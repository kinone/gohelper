package pool

import (
	"io/ioutil"
	"log"
	"time"
)

type StdLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type ProcessPool struct {
	pool   chan struct{}
	Logger StdLogger
}

func NewProcessPool(num int) *ProcessPool {
	pool := make(chan struct{}, num)
	logger := log.New(ioutil.Discard, "", log.LstdFlags)

	return &ProcessPool{
		pool:   pool,
		Logger: logger,
	}
}

func (p *ProcessPool) Go(callable func()) {
	p.add()
	go func() {
		defer p.done()
		callable()
	}()
}

func (p *ProcessPool) Close() {
	p.Logger.Println("Closing ProcessPool")
	close(p.pool)
	for !p.empty() { // 等待所有协程执行完成
		time.Sleep(time.Microsecond * 10)
	}
	p.Logger.Println("ProcessPool Closed")
}

func (p *ProcessPool) empty() bool {
	return len(p.pool) == 0
}

func (p *ProcessPool) add() {
	p.pool <- struct{}{}
}

func (p *ProcessPool) done() {
	<-p.pool
}
