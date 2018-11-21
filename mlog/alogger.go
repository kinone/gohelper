package mlog

type AsyncLogger struct {
	*Logger
}

func NewAsyncLogger(file string, strLevel string, async bool) *AsyncLogger {
	l := NewLogger(file, strLevel)

	if async {
		h := newSmartHandler(l.h)
		go h.start()

		l.h = h
	}

	return &AsyncLogger{l}
}
