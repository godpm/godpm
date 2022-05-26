package process

import (
	"path/filepath"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

// newLog
func (p *Process) newLog(path string, maxBytes int, backups int) *lumberjack.Logger {

	return &lumberjack.Logger{
		Filename:   fixPath(*p.conf.Directory, path),
		MaxSize:    maxBytes,
		MaxBackups: backups,
	}
}

func fixPath(dir, path string) string {
	if len(dir) == 0 {
		return path
	}

	if strings.HasPrefix(path, dir) {
		return path
	}

	return filepath.Join(dir, path)
}
