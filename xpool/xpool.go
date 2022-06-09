package xpool

import (
	"sync"
)

// Xpool like sync.WaitGroup, but it can control
// the number of goroutines concurrent running.
type Xpool struct {
	q chan struct{}
	g *sync.WaitGroup
}

func New(size int) *Xpool {
	if size < 1 {
		size = 1
	}
	return &Xpool{
		q: make(chan struct{}, size),
		g: &sync.WaitGroup{},
	}
}

// Add the number of goroutines you ready yo run.
// It may block when too many goroutines are running.
//
// 'delta' should not less than 1.
// delta=1 is suggested in case of unnecessary blocking.
func (p *Xpool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.q <- struct{}{}
		p.g.Add(1)
	}
}

func (p *Xpool) Done() {
	<-p.q
	p.g.Done()
}

func (p *Xpool) Wait() {
	p.g.Wait()
}
