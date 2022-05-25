package process

// TODO
/*
func (p *Process) dir() string {
	if p.conf.Directory != nil {
		return *p.conf.Directory
	}

	return ""
}*/

func (p *Process) stdoutMaxBytes() int64 {
	if p.conf.StdoutMaxBytes > 0 {
		return p.conf.StdoutMaxBytes
	}

	// default  100M
	return 100 * 1024
}

func (p *Process) stdoutMaxBackups() int {
	if p.conf.StdoutMaxBackups > 0 {
		return p.conf.StdoutMaxBackups
	}

	return 10
}

func (p *Process) stderrMaxBytes() int64 {
	if p.conf.StderrMaxBytes > 0 {
		return p.conf.StderrMaxBytes
	}

	// defautl 100
	return 100 * 1024
}

func (p *Process) stderrMaxBackups() int {
	if p.conf.StderrMaxBackups > 0 {
		return p.conf.StderrMaxBackups
	}

	return 10
}
