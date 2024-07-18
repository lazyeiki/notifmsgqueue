package notifmsgqueue

import (
	"fmt"
	"runtime"
	"sync"
)

var ErrQueueFull = fmt.Errorf("queue is full")

type MsgQueueMode int

const (
	Single MsgQueueMode = iota
	WorkerPool
)

type NotifMsgQueue struct {
	mu      sync.Mutex
	queue   chan any
	running bool
	mode    MsgQueueMode
	workers int
	wg      sync.WaitGroup
}

func New(bufferSize int, mode MsgQueueMode, workers int) *NotifMsgQueue {
	if mode == Single {
		workers = 1
	} else {
		numCPU := runtime.NumCPU()
		if workers <= 0 || workers > numCPU {
			workers = numCPU
		}
	}
	return &NotifMsgQueue{
		queue:   make(chan any, bufferSize),
		mode:    mode,
		workers: workers,
	}
}

func (p *NotifMsgQueue) Push(msg any) error {
	select {
	case p.queue <- msg:
		return nil
	default:
		return ErrQueueFull
	}
}

func (p *NotifMsgQueue) worker(h func(f any)) {
	defer p.wg.Done()
	for msg := range p.queue {
		h(msg)
	}
}

func (p *NotifMsgQueue) Run(h func(f any)) {
	p.mu.Lock()
	if p.running {
		p.mu.Unlock()
		return
	}
	p.running = true
	p.mu.Unlock()

	if p.mode == WorkerPool {
		for i := 0; i < p.workers; i++ {
			p.wg.Add(1)
			go p.worker(h)
		}
	} else {
		go func() {
			for msg := range p.queue {
				h(msg)
			}
		}()
	}
}

func (p *NotifMsgQueue) Stop() {
	p.mu.Lock()
	if !p.running {
		p.mu.Unlock()
		return
	}
	p.running = false
	close(p.queue)
	p.mu.Unlock()

	if p.mode == WorkerPool {
		p.wg.Wait()
	}
}
