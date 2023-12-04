package logging

import (
	"fmt"
	"log"
)

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLogger) Tracef(template string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(template, vals...))
}
