package process

const (
	// StateRunning process state running
	StateRunning = "running"
	// StateStarting process state starting
	StateStarting = "starting"
	// StateStopped process state stopped
	StateStopped = "stopped"
	// StateStopping process state stopping, state between stopped and running
	StateStopping = "stopping"
	// StateExited process exited
	StateExited = "exited"
	// StateFatal process fatal
	StateFatal = "fatal"
	// StateBackOff
	StateBackOff = "back_off"
)

func (p *Process) changeStateTo(state string) {
	p.lock.Lock()
	p.state = state
	p.lock.Unlock()
}

func (p *Process) stateIs(state string) bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.state == state
}

// State state
func (p *Process) State() string {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.state
}
