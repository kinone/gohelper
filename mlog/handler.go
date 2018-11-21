package mlog

import (
	"log"
	"os"
	"sync"
)

type Handler interface {
	Log(level int, msg string, context ...interface{})
	Logf(level int, format string, context ...interface{})
	Reload() error
	Close()
}

type filter func(level int) bool

type fileHandler struct {
	logger *log.Logger
	fname  string
	file   *os.File
	pass   filter
}

func newFileHandler(fname string, pass filter) *fileHandler {
	var (
		f   *os.File
		l   *log.Logger
		err error
	)

	if len(fname) == 0 {
		l = log.New(os.Stdout, "", log.LstdFlags)
	} else {
		f, err = os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if nil != err {
			log.Panic("can not open log file", err)
		}

		l = log.New(f, "", log.LstdFlags)
	}

	return &fileHandler{l, fname, f, pass}
}

func (h *fileHandler) Log(level int, msg string, context ...interface{}) {
	lp, exist := levelPrefix[level]
	if !exist { // 不存在的level
		return
	}

	if h.pass(level) {
		return
	}

	context = append([]interface{}{lp, msg}, context...)
	h.logger.Println(context...)
}

func (h *fileHandler) Logf(level int, format string, context ...interface{}) {
	lp, exist := levelPrefix[level]
	if !exist { // 不存在的level
		return
	}

	if h.pass(level) {
		return
	}

	context = append([]interface{}{lp}, context...)
	h.logger.Printf("%s "+format, context...)
}
func (h *fileHandler) Reload() error {
	if nil == h.file {
		return nil
	}

	lf, err := os.OpenFile(h.fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		return err
	}
	h.logger.SetOutput(lf)
	h.file.Close()
	h.file = lf

	return nil
}

func (h *fileHandler) Close() {
	if nil != h.file {
		h.file.Close()
	}
}

type smartHandler struct {
	Handler
	c    chan func()
	quit chan struct{}
	wg   *sync.WaitGroup
}

func newSmartHandler(fh Handler) *smartHandler {
	c := make(chan func())
	quit := make(chan struct{})
	wg := new(sync.WaitGroup)

	return &smartHandler{fh, c, quit, wg}
}

func (h *smartHandler) Logf(level int, format string, context ...interface{}) {
	h.c <- func() {
		h.Handler.Logf(level, format, context...)
	}
}

func (h *smartHandler) Log(level int, msg string, context ...interface{}) {
	h.c <- func() {
		h.Handler.Log(level, msg, context...)
	}
}

func (h *smartHandler) start() {
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		h.Handler.Log(LDEBUG, "AysncLogger started")
		defer h.Handler.Log(LDEBUG, "AsyncLogger stoped")

		for {
			select {
			case f := <-h.c:
				f()
			case <-h.quit:
				return
			}
		}
	}()
}

func (h *smartHandler) Close() {
	defer h.Handler.Close()

	h.quit <- struct{}{}
	h.wg.Wait()
}
