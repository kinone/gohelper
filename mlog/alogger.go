package mlog

import (
	"sync"
)

type AsyncLogger struct {
	Logger
	async bool
	c     chan func()
	quit  chan struct{}
}

var wg sync.WaitGroup

func NewAsyncLogger(file string, strLevel string, async bool) *AsyncLogger {
	l := NewLogger(file, strLevel)
	ch := make(chan func())
	quit := make(chan struct{})
	al := &AsyncLogger{*l, async, ch, quit}

	if async {
		al.start()
	}

	return al
}

func (l *AsyncLogger) Debug(msg string, context ...interface{}) {
	l.Log(LDEBUG, msg, context...)
}

func (l *AsyncLogger) Debugf(format string, context ...interface{}) {
	l.Logf(LDEBUG, format, context...)
}

func (l *AsyncLogger) Info(msg string, context ...interface{}) {
	l.Log(LINFO, msg, context...)
}

func (l *AsyncLogger) Infof(format string, context ...interface{}) {
	l.Logf(LINFO, format, context...)
}

func (l *AsyncLogger) Notice(msg string, context ...interface{}) {
	l.Log(LNOTICE, msg, context...)
}

func (l *AsyncLogger) Noticef(format string, context ...interface{}) {
	l.Logf(LNOTICE, format, context...)
}

func (l *AsyncLogger) Warning(msg string, context ...interface{}) {
	l.Log(LWARNING, msg, context...)
}

func (l *AsyncLogger) Warningf(format string, context ...interface{}) {
	l.Logf(LWARNING, format, context...)
}

func (l *AsyncLogger) Error(msg string, context ...interface{}) {
	l.Log(LERROR, msg, context...)
}

func (l *AsyncLogger) Errorf(format string, context ...interface{}) {
	l.Logf(LERROR, format, context...)
}

func (l *AsyncLogger) Critical(msg string, context ...interface{}) {
	l.Log(LCRITICAL, msg, context...)
}

func (l *AsyncLogger) Criticalf(format string, context ...interface{}) {
	l.Logf(LCRITICAL, format, context...)
}

func (l *AsyncLogger) Alert(msg string, context ...interface{}) {
	l.Log(LALERT, msg, context...)
}

func (l *AsyncLogger) Alertf(format string, context ...interface{}) {
	l.Logf(LALERT, format, context...)
}

func (l *AsyncLogger) Emergency(msg string, context ...interface{}) {
	l.Log(LEMERGENCY, msg, context...)
}

func (l *AsyncLogger) Emergencyf(format string, context ...interface{}) {
	l.Logf(LEMERGENCY, format, context...)
}

func (l *AsyncLogger) Log(level int, msg string, context ...interface{}) {
	if l.async {
		l.c <- func() {
			l.Logger.Log(level, msg, context...)
		}
	} else {
		l.Logger.Log(level, msg, context...)
	}
}

func (l *AsyncLogger) Logf(level int, format string, context ...interface{}) {
	if l.async {
		l.c <- func() {
			l.Logger.Logf(level, format, context...)
		}
	} else {
		l.Logger.Logf(level, format, context...)
	}
}

func (l *AsyncLogger) Close() {
	if l.async {
		l.quit <- struct{}{}
		wg.Wait()
	}

	l.Logger.Close()
}

func (l *AsyncLogger) start() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		l.Logger.Debug("AysncLogger started")
		for {
			select {
			case f := <-l.c:
				f()
			case <-l.quit:
				l.Logger.Debug("AsyncLogger stoped")
				return
			}
		}
	}()
}
