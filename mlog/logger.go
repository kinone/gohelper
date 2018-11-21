package mlog

import (
	"log"
)

type Logger struct {
	*log.Logger
	h     Handler
	level int
}

func NewLogger(file string, strLevel string) *Logger {
	level := ConvertLogLevel(strLevel)

	h := newFileHandler(file, func(l int) bool {
		return l < level
	})

	return &Logger{h.logger, h, level}
}

func (l *Logger) Debug(msg string, context ...interface{}) {
	l.h.Log(LDEBUG, msg, context...)
}

func (l *Logger) Debugf(format string, context ...interface{}) {
	l.h.Logf(LDEBUG, format, context...)
}

func (l *Logger) Info(msg string, context ...interface{}) {
	l.h.Log(LINFO, msg, context...)
}

func (l *Logger) Infof(format string, context ...interface{}) {
	l.h.Logf(LINFO, format, context...)
}

func (l *Logger) Notice(msg string, context ...interface{}) {
	l.h.Log(LNOTICE, msg, context...)
}

func (l *Logger) Noticef(format string, context ...interface{}) {
	l.h.Logf(LNOTICE, format, context...)
}

func (l *Logger) Warning(msg string, context ...interface{}) {
	l.h.Log(LWARNING, msg, context...)
}

func (l *Logger) Warningf(format string, context ...interface{}) {
	l.h.Logf(LWARNING, format, context...)
}

func (l *Logger) Error(msg string, context ...interface{}) {
	l.h.Log(LERROR, msg, context...)
}

func (l *Logger) Errorf(format string, context ...interface{}) {
	l.h.Logf(LERROR, format, context...)
}

func (l *Logger) Critical(msg string, context ...interface{}) {
	l.h.Log(LCRITICAL, msg, context...)
}

func (l *Logger) Criticalf(format string, context ...interface{}) {
	l.h.Logf(LCRITICAL, format, context...)
}

func (l *Logger) Alert(msg string, context ...interface{}) {
	l.h.Log(LALERT, msg, context...)
}

func (l *Logger) Alertf(format string, context ...interface{}) {
	l.h.Logf(LALERT, format, context...)
}

func (l *Logger) Emergency(msg string, context ...interface{}) {
	l.h.Log(LEMERGENCY, msg, context...)
}

func (l *Logger) Emergencyf(format string, context ...interface{}) {
	l.h.Logf(LEMERGENCY, format, context...)
}

func (l *Logger) Reload() error {
	return l.h.Reload()
}

func (l *Logger) Close() {
	l.h.Close()
}
