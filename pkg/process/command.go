package process

import (
	"os"
	"os/exec"
	"os/user"
	//	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/godpm/godpm/pkg/log"
)

func (p *Process) createCommand() (err error) {
	args := parseCommand(p.conf.Command)
	p.cmd = exec.Command(args[0])
	if len(args) > 1 {
		p.cmd.Args = args
	}

	p.cmd.SysProcAttr = &syscall.SysProcAttr{}
	p.setEnv()
	p.setLog()
	p.setDirectory()
	err = p.setUser(p.conf.User)
	if err != nil {
		log.Error().Printf("process %s set user failed %#v", p.conf.Name, err)
		return
	}

	// signal can send to process's child process
	setpgid(p.cmd.SysProcAttr)
	return
}

func setpgid(attr *syscall.SysProcAttr) {
	attr.Setpgid = true
}

func parseCommand(command string) (args []string) {
	splits := strings.Split(command, " ")
	args = make([]string, 0, len(splits))
	for _, v := range splits {
		if len(v) > 0 {
			args = append(args, v)
		}
	}

	return args
}

func (p *Process) setUser(username string) (err error) {
	if len(username) == 0 {
		return
	}

	u, err := user.Lookup(username)
	if err != nil {
		return
	}

	uid, err := strconv.ParseUint(u.Uid, 10, 32)
	if err != nil {
		return
	}

	gid, err := strconv.ParseUint(u.Gid, 10, 32)
	if err != nil {
		return
	}

	p.cmd.SysProcAttr.Credential = &syscall.Credential{
		Gid: uint32(gid),
		Uid: uint32(uid),
	}

	return
}

func (p *Process) setEnv() {
	confEnv := strings.Split(p.conf.Environment, ";")
	env := append(os.Environ(), confEnv...)
	p.cmd.Env = append(p.cmd.Env, env...)
}

func (p *Process) setDirectory() {
	if p.conf.Directory != nil {
		p.cmd.Dir = *p.conf.Directory
	}
}

func (p *Process) setLog() {
	if p.conf.StdoutFile != nil {
		p.cmd.Stdout = p.newLog(*p.conf.StdoutFile, p.stdoutMaxBytes(), p.stdoutMaxBackups())
	}

	if p.conf.StderrFile != nil {
		p.cmd.Stderr = p.newLog(*p.conf.StderrFile, p.stderrMaxBytes(), p.stderrMaxBackups())
	}
}
