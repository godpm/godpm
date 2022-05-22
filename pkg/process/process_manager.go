package process

import (
	"fmt"
	"sync"

	"github.com/godpm/godpm/config"
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
		pm.procs[proc.name()] = proc
	}

	return proc
}

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
func (pm *Manager) Start(name string) (err error) {
	proc, err := pm.findOrError(name)
	if err != nil {
		return
	}

	return proc.Start()
}

// Restart restart a process
func (pm *Manager) Restart(name string) (err error) {
	proc, err := pm.findOrError(name)
	if err != nil {
		return
	}

	err = proc.Stop()
	if err != nil {
		return
	}

	return proc.Stop()
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
			err := proc.Start()
			if err != nil {
				return false
			}
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

// CreateProcess create process
func CreateProcess(conf *config.ProcessConfig) *Process {
	return pm.CreateProcess(conf)
}
