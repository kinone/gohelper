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
	pool chan struct{}
}

var Logger StdLogger

func init() {
	Logger = log.New(ioutil.Discard, "", log.LstdFlags)
}

func NewProcessPool(num int) *ProcessPool {
	pool := make(chan struct{}, num)

	return &ProcessPool{
		pool: pool,
	}
}

func (p *ProcessPool) Go(callable func(...interface{}), params ...interface{}) {
	p.add()
	go func() {
		defer p.done()
		callable(params...)
	}()
}

func (p *ProcessPool) Close() {
	Logger.Println("Closing ProcessPool")
	close(p.pool)
	for !p.empty() { // 等待所有协程执行完成
		time.Sleep(time.Microsecond * 10)
	}
	Logger.Println("ProcessPool Closed")
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
