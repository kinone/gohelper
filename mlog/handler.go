package mlog

import (
	"log"
	"os"
	"sync"
)

type Handler interface {
	Log(level int, context ...interface{})
	Logf(level int, format string, context ...interface{})
	Reload() error
	Close()
}

type NullHandler struct {
}

func NewNullHandler() *NullHandler {
	return &NullHandler{}
}

func (h *NullHandler) Log(int, ...interface{})          {}
func (h *NullHandler) Logf(int, string, ...interface{}) {}
func (h *NullHandler) Reload() error                    { return nil }
func (h *NullHandler) Close()                           {}

type Filter func(level int) bool

type FileHandler struct {
	logger *log.Logger
	fname  string
	file   *os.File
	filter Filter
}

func NewFileHandler(fname string, filter Filter) *FileHandler {
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

	return &FileHandler{l, fname, f, filter}
}

func (h *FileHandler) Log(level int, context ...interface{}) {
	lp, exist := levelPrefix[level]
	if !exist { // 不存在的level
		return
	}

	if !h.filter(level) {
		return
	}

	context = append([]interface{}{lp}, context...)
	h.logger.Println(context...)
}

func (h *FileHandler) Logf(level int, format string, context ...interface{}) {
	lp, exist := levelPrefix[level]
	if !exist { // 不存在的level
		return
	}

	if !h.filter(level) {
		return
	}

	context = append([]interface{}{lp}, context...)
	h.logger.Printf("%s "+format, context...)
}
func (h *FileHandler) Reload() error {
	if nil == h.file {
		return nil
	}

	f, err := os.OpenFile(h.fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		return err
	}
	h.logger.SetOutput(f)
	h.file.Close()
	h.file = f

	return nil
}

func (h *FileHandler) Close() {
	if nil != h.file {
		h.file.Close()
	}
}

type SmartHandler struct {
	Handler
	ch   chan func()
	quit chan struct{}
	wg   *sync.WaitGroup
}

func NewSmartHandler(fh Handler) *SmartHandler {
	c := make(chan func())
	quit := make(chan struct{})
	wg := new(sync.WaitGroup)

	return &SmartHandler{fh, c, quit, wg}
}

func (h *SmartHandler) Log(level int, context ...interface{}) {
	h.ch <- func() {
		h.Handler.Log(level, context...)
	}
}

func (h *SmartHandler) Logf(level int, format string, context ...interface{}) {
	h.ch <- func() {
		h.Handler.Logf(level, format, context...)
	}
}

func (h *SmartHandler) start() {
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		h.Handler.Log(LDEBUG, "AsyncLogger started")
		defer h.Handler.Log(LDEBUG, "AsyncLogger stoped")

		for {
			select {
			case f := <-h.ch:
				f()
			case <-h.quit:
				return
			}
		}
	}()
}

func (h *SmartHandler) Close() {
	defer h.Handler.Close()

	h.quit <- struct{}{}
	h.wg.Wait()
}
