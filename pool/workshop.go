package pool

import (
	"sync"
)

type Workshop struct {
	jobs chan func()
	done []chan struct{}
	wg   *sync.WaitGroup
}

func NewWorkshop(workerNum int) *Workshop {
	jobs := make(chan func())
	done := make([]chan struct{}, workerNum)
	wg := new(sync.WaitGroup)

	ws := &Workshop{jobs, done, wg}

	for i := 0; i < workerNum; i++ {
		done[i] = make(chan struct{})
		go ws.worker(i)
	}

	return ws
}

func (ws *Workshop) Do(callable func()) {
	ws.jobs <- callable
}

func (ws *Workshop) Close() {
	for _, done := range ws.done {
		done <- struct{}{}
	}

	ws.wg.Wait()
}

func (ws *Workshop) worker(idx int) {
	ws.wg.Add(1)
	defer ws.wg.Done()
	Logger.Printf("Worker %d started", idx)

	for {
		select {
		case f := <-ws.jobs:
			f()
		case <-ws.done[idx]:
			Logger.Printf("Worker %d stoped", idx)
			return
		}
	}
}
