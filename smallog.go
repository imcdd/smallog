//go:build !windows && !plan9

package smallog

import (
	"errors"
	"log/syslog"
)

type Logger interface {
	Emerg(string, ...interface{})
	Alert(string, ...interface{})
	Crit(string, ...interface{})
	Err(string, ...interface{})
	Warning(string, ...interface{})
	Notice(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

var logLevel = syslog.LOG_INFO
var logger Logger = nil

func init() {
	sysLogger, _ := syslog.New(syslog.LOG_DEBUG, "")
	if err != nil {
		panic("failed to new syslog ")
	}
	logger = newDefaultLogger(sysLogger)
	SetLogLevel(syslog.LOG_INFO)
}

func SetCustomLogger(newLogger Logger) (err error) {
	if newLogger == nil {
		err = errors.New("new logger is nil")
	}
	logger = newLogger
	SetLogLevel(logLevel)
	return
}

func GetLogLevel() (level syslog.Priority) {
	level = logLevel
	return
}

func SetLogLevel(level syslog.Priority) {

	{
		debug = nothingToDo
		info = nothingToDo
		notice = nothingToDo
		warning = nothingToDo
		err = nothingToDo
		crit = nothingToDo
		alert = nothingToDo
		emerg = nothingToDo
	}

resetDefault:
	switch level {
	case syslog.LOG_DEBUG:
		debug = logger.Debug
		fallthrough
	case syslog.LOG_INFO:
		info = logger.Info
		fallthrough
	case syslog.LOG_NOTICE:
		notice = logger.Notice
		fallthrough
	case syslog.LOG_WARNING:
		warning = logger.Warning
		fallthrough
	case syslog.LOG_ERR:
		err = logger.Err
		fallthrough
	case syslog.LOG_CRIT:
		crit = logger.Crit
		fallthrough
	case syslog.LOG_ALERT:
		alert = logger.Alert
		fallthrough
	case syslog.LOG_EMERG:
		emerg = logger.Emerg
		logLevel = level
		return
	default:
		logger.Err("failed to set smallog level, fallback to default level [LOG_INFO]")
		level = syslog.LOG_INFO
		goto resetDefault
	}
}

func nothingToDo(string, ...interface{}) {}

var (
	debug   func(string, ...interface{})
	info    func(string, ...interface{})
	notice  func(string, ...interface{})
	warning func(string, ...interface{})
	err     func(string, ...interface{})
	crit    func(string, ...interface{})
	alert   func(string, ...interface{})
	emerg   func(string, ...interface{})
)

func Debug(format string, a ...interface{}) {
	debug(format, a...)
}

func Info(format string, a ...interface{}) {
	info(format, a...)
}

func Notice(format string, a ...interface{}) {
	notice(format, a...)
}

func Warning(format string, a ...interface{}) {
	warning(format, a...)
}

func Err(format string, a ...interface{}) {
	err(format, a...)
}

func Crit(format string, a ...interface{}) {
	crit(format, a...)
}

func Alert(format string, a ...interface{}) {
	alert(format, a...)
}

func Emerg(format string, a ...interface{}) {
	emerg(format, a...)
}
