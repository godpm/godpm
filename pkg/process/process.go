package process

import (
	"errors"
	"os/exec"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/godpm/godpm/config"
	"github.com/godpm/godpm/pkg/log"
)

// Process ..
type Process struct {
	startAt       time.Time
	stopAt        time.Time
	state         string
	conf          *config.ProcessConfig
	cmd           *exec.Cmd
	lock          sync.RWMutex
	retryTimes    int
	manualStopped bool
}

// New new process
func New(conf *config.ProcessConfig) *Process {
	retry := 3
	if conf.RetryTimes != nil {
		retry = *conf.RetryTimes
	}

	return &Process{
		cmd:        nil,
		conf:       conf,
		state:      StateStopped,
		retryTimes: retry,
	}
}

// Name process name
func (p *Process) Name() string {
	return p.conf.Name
}

func (p *Process) getStartSecs() int {
	if p.conf.StartSecs != nil {
		return *p.conf.StartSecs
	}

	return 1
}

func (p *Process) getRestartPauseSecs() int {
	if p.conf.RestartPause != nil {
		return *p.conf.RestartPause
	}

	return 1
}

// 	State start the process
func (p *Process) Start(wait bool) (err error) {
	log.Info().Printf("try to start program: %s", p.Name())

	err = p.createCommand()
	if err != nil {
		log.Error().Printf("try to start program: %s failed %#v", p.Name(), err)
		return
	}

	retryTimes := 0
	startSecs := p.getStartSecs()
	restartPauseSecs := p.getRestartPauseSecs()
	finishChan := make(chan struct{})

	once := sync.Once{}
	waitFunc := func() {
		once.Do(func() {
			if wait {
				finishChan <- struct{}{}
			}
		})
	}

	go func() {
		for !p.manualStopped {
			if retryTimes > p.retryTimes {
				log.Error().Printf("failed to start program: %s, because try time is greater than %d", p.Name(), p.retryTimes)
				p.changeStateTo(StateFatal)
				break
			}

			if restartPauseSecs > 0 && retryTimes != 0 {
				time.Sleep(time.Duration(restartPauseSecs) * time.Second)
			}

			retryTimes++
			p.changeStateTo(StateStarting)
			p.startAt = time.Now()
			err = p.cmd.Start()
			if err != nil {
				log.Error().Printf("failed to start program: %s, err: %#v", p.Name(), err)
				p.changeStateTo(StateBackOff)
				waitFunc()
				continue
			}

			var (
				processExited = int32(0)
			)

			if startSecs <= 0 {
				p.changeStateTo(StateRunning)
			} else {
				go p.checkIfProgramIsRunning(time.Duration(startSecs)*time.Second, &processExited)
			}
			waitFunc()

			p.waitForExist()

			atomic.StoreInt32(&processExited, 1)

			if p.stateIs(StateRunning) {
				p.changeStateTo(StateExited)
			} else {
				p.changeStateTo(StateBackOff)
			}

			if !p.conf.AutoRestart {
				break
			}
		}

		p.manualStopped = false
	}()

	if wait {
		<-finishChan
	}

	return
}

// Pid pid info
func (p *Process) Pid() int {
	if p.cmd != nil && p.cmd.Process != nil {
		return p.cmd.Process.Pid
	}

	return -1
}

// Uptime uptime
func (p *Process) Uptime() time.Time {
	return p.startAt
}

// checkIfProgramIsRunning wait untile endTime and check if program is starting
func (p *Process) checkIfProgramIsRunning(duration time.Duration, processExited *int32) {
	end := time.Now().Add(duration)
	for atomic.LoadInt32(processExited) == 0 && time.Now().Before(end) {
		time.Sleep(100 * time.Millisecond)
	}

	if p.state == StateStarting && atomic.LoadInt32(processExited) == 0 {
		p.changeStateTo(StateRunning)
	}
}

func (p *Process) waitForExist() {
	if err := p.cmd.Wait(); err != nil {
		log.Error().Println("command wait failed ", err)
		return
	}

	p.stopAt = time.Now()
}

// Stop stop process
func (p *Process) Stop() (err error) {
	p.lock.Lock()
	p.manualStopped = true
	p.lock.Unlock()

	if !p.stateIs(StateRunning) {
		log.Error().Println("program is not running")
		err = errors.New("program is not running")
		return
	}

	p.changeStateTo(StateStopping)
	err = p.kill()
	if err != nil {
		log.Error().Printf("stop program: %s failed: %#v ", p.Name(), err)
		return
	}

	p.changeStateTo(StateExited)
	return
}

func (p *Process) kill() (err error) {

	// include child process
	pid := -(p.cmd.Process.Pid)
	return syscall.Kill(pid, syscall.SIGKILL)
}
