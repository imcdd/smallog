//go:build !windows && !plan9

package smallog

import (
	"fmt"
	"log/syslog"
)

type defaultLogger struct {
	logger *syslog.Writer
}

func (d *defaultLogger) Emerg(s string, i ...interface{}) {
	_ = d.logger.Emerg(fmt.Sprintf(s, i))
}

func (d *defaultLogger) Alert(s string, i ...interface{}) {
	_ = d.logger.Alert(fmt.Sprintf(s, i))
}

func (d *defaultLogger) Crit(s string, i ...interface{}) {
	_ = d.logger.Crit(fmt.Sprintf(s, i))
}

func (d *defaultLogger) Err(s string, i ...interface{}) {
	_ = d.logger.Err(fmt.Sprintf(s, i))
}

func (d *defaultLogger) Warning(s string, i ...interface{}) {
	_ = d.logger.Warning(fmt.Sprintf(s, i))
}

func (d *defaultLogger) Notice(s string, i ...interface{}) {
	_ = d.logger.Notice(fmt.Sprintf(s, i))
}

func (d *defaultLogger) Info(s string, i ...interface{}) {
	_ = d.logger.Info(fmt.Sprintf(s, i))
}

func (d *defaultLogger) Debug(s string, i ...interface{}) {
	_ = d.logger.Debug(fmt.Sprintf(s, i))
}

func newDefaultLogger(logger *syslog.Writer) (dl *defaultLogger) {
	dl = &defaultLogger{logger: logger}
	return
}
