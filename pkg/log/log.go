package log

import (
	"log"
	"os"
)

const (
	debugLevel = "DEBUG "
	infoLevel  = "INFO "
	errorLevel = "ERROR "
	fatalLevel = "FATAL "
	panicLevel = "PANIC "

	stdFlag = log.Llongfile | log.LstdFlags
)

// Info new info prefix log
func Info() *log.Logger {
	return log.New(os.Stdout, infoLevel, stdFlag)
}

// Error new error prefix log
func Error() *log.Logger {
	return log.New(os.Stderr, errorLevel, stdFlag)
}

// Fatal new fatal prefix log
func Fatal() *log.Logger {
	return log.New(os.Stderr, fatalLevel, stdFlag)
}

// Panic new panic prefix log
func Panic() *log.Logger {
	return log.New(os.Stderr, panicLevel, stdFlag)
}

// Debug new debug prefix log
func Debug() *log.Logger {
	return log.New(os.Stdout, debugLevel, stdFlag)
}
