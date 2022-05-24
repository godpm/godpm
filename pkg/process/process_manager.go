package process

import (
	"fmt"
	"sync"

	"github.com/godpm/godpm/config"
	"github.com/godpm/godpm/pkg/log"
)

// Manager ..
type Manager struct {
	procs map[string]*Process
	lock  sync.Mutex
}

var pm *Manager

func init() {
	pm = &Manager{
		procs: make(map[string]*Process),
		lock:  sync.Mutex{},
	}
}

// CreateProcess create process
func (pm *Manager) CreateProcess(conf *config.ProcessConfig) *Process {
	pm.lock.Lock()
	defer pm.lock.Unlock()

	proc, ok := pm.procs[conf.Name]
	if !ok {
		proc := New(conf)
		pm.procs[proc.Name()] = proc
	}

	return proc
}

//
func (pm *Manager) Reread() {}

// Find find one process
func (pm *Manager) Find(name string) (proc *Process, ok bool) {
	pm.lock.Lock()
	proc, ok = pm.procs[name]
	pm.lock.Unlock()
	return
}

func (pm *Manager) findOrError(name string) (proc *Process, err error) {
	proc, ok := pm.Find(name)
	if !ok {
		err = fmt.Errorf("process %s not found", name)
	}

	return
}

// Stop stop a process
func (pm *Manager) Stop(name string) (err error) {
	proc, err := pm.findOrError(name)
	if err != nil {
		return
	}

	return proc.Stop()
}

// Start start a process
func (pm *Manager) Start(name string, wait bool) (err error) {
	proc, err := pm.findOrError(name)
	if err != nil {
		return
	}

	return proc.Start(wait)
}

// Restart restart a process
func (pm *Manager) Restart(name string) (err error) {
	proc, err := pm.findOrError(name)
	if err != nil {
		return
	}

	_ = proc.Stop()
	return proc.Start(false)
}

// Remove remove a process
// TODO
func (pm *Manager) Remove(name string) (err error) {
	return
}

// StartAutoStart all the auto start process
func (pm *Manager) StartAutoStart() (err error) {
	pm.Range(func(proc *Process) bool {
		if proc.conf.AutoStart {
			go func() {
				err := proc.Start(false)
				if err != nil {
					log.Error().Println("start failed", err)
				}
			}()
		}

		return true
	})

	return
}

// Range range current process
func (pm *Manager) Range(f func(proc *Process) bool) {
	pm.lock.Lock()
	defer pm.lock.Unlock()

	for _, proc := range pm.procs {
		if !f(proc) {
			break
		}
	}
}

// List list all process
func (pm *Manager) List() (procs []*Process) {
	procs = make([]*Process, 0, len(pm.procs))

	pm.Range(func(proc *Process) bool {
		procs = append(procs, proc)
		return true
	})

	return
}

// CreateProcess create process
func CreateProcess(conf *config.ProcessConfig) *Process {
	return pm.CreateProcess(conf)
}

// StartAutoStartProcesses start all the autostart process
func StartAutoStartProcesses() (err error) {
	return pm.StartAutoStart()
}

// Find find a process
func Find(name string) (*Process, bool) {
	return pm.Find(name)
}

func Restart(name string) (err error) {
	return pm.Restart(name)
}

// Start start a process
func Start(name string, wait bool) error {
	return pm.Start(name, wait)
}

// Stop stop a process
func Stop(name string) error {
	return pm.Stop(name)
}

// List list all process
func List() []*Process {
	return pm.List()
}

// InitAndStart init process and start all the autostart process
func InitAndStart() {
	for _, pc := range config.AppConfig.GetAllProcesssConfig() {
		pm.CreateProcess(pc)
	}

	err := pm.StartAutoStart()
	if err != nil {
		log.Error().Println("start failed ", err)
	}
}

// Reread reread process config
func Reread() {
	pm.Reread()
}
