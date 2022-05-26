package process

// TODO
/*
func (p *Process) dir() string {
	if p.conf.Directory != nil {
		return *p.conf.Directory
	}

	return ""
}*/

func (p *Process) stdoutMaxBytes() int {
	if p.conf.StdoutMaxSize > 0 {
		return p.conf.StdoutMaxSize
	}

	// default  100MB
	return 100
}

func (p *Process) stdoutMaxBackups() int {
	if p.conf.StdoutMaxBackups > 0 {
		return p.conf.StdoutMaxBackups
	}

	return 10
}

func (p *Process) stderrMaxBytes() int {
	if p.conf.StderrMaxSize > 0 {
		return p.conf.StderrMaxSize
	}

	// defautl 100MB
	return 100
}

func (p *Process) stderrMaxBackups() int {
	if p.conf.StderrMaxBackups > 0 {
		return p.conf.StderrMaxBackups
	}

	return 10
}
