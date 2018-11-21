package mlog

type Logger struct {
	h     Handler
	level int
}

func NewLogger(file string, strLevel string) *Logger {
	level := ConvertLogLevel(strLevel)

	h := NewFileHandler(file, func(l int) bool {
		return l >= level
	})

	return &Logger{h, level}
}

func (l *Logger) Print(v ...interface{}) {
	l.h.Log(NONE, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.h.Log(NONE, v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.h.Logf(NONE, format, v...)
}

func (l *Logger) Debug(context ...interface{}) {
	l.h.Log(LDEBUG, context...)
}

func (l *Logger) Debugf(format string, context ...interface{}) {
	l.h.Logf(LDEBUG, format, context...)
}

func (l *Logger) Info(context ...interface{}) {
	l.h.Log(LINFO, context...)
}

func (l *Logger) Infof(format string, context ...interface{}) {
	l.h.Logf(LINFO, format, context...)
}

func (l *Logger) Notice(context ...interface{}) {
	l.h.Log(LNOTICE, context...)
}

func (l *Logger) Noticef(format string, context ...interface{}) {
	l.h.Logf(LNOTICE, format, context...)
}

func (l *Logger) Warning(context ...interface{}) {
	l.h.Log(LWARNING, context...)
}

func (l *Logger) Warningf(format string, context ...interface{}) {
	l.h.Logf(LWARNING, format, context...)
}

func (l *Logger) Error(context ...interface{}) {
	l.h.Log(LERROR, context...)
}

func (l *Logger) Errorf(format string, context ...interface{}) {
	l.h.Logf(LERROR, format, context...)
}

func (l *Logger) Critical(context ...interface{}) {
	l.h.Log(LCRITICAL, context...)
}

func (l *Logger) Criticalf(format string, context ...interface{}) {
	l.h.Logf(LCRITICAL, format, context...)
}

func (l *Logger) Alert(context ...interface{}) {
	l.h.Log(LALERT, context...)
}

func (l *Logger) Alertf(format string, context ...interface{}) {
	l.h.Logf(LALERT, format, context...)
}

func (l *Logger) Emergency(context ...interface{}) {
	l.h.Log(LEMERGENCY, context...)
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

type AsyncLogger struct {
	*Logger
}

func NewAsyncLogger(file string, strLevel string, async bool) *AsyncLogger {
	l := NewLogger(file, strLevel)

	if async {
		h := NewSmartHandler(l.h)
		go h.start()

		l.h = h
	}

	return &AsyncLogger{l}
}

type NullLogger struct {
	*Logger
}

func NewNullLogger() *NullLogger {
	return &NullLogger{&Logger{NewNullHandler(), NONE}}
}
