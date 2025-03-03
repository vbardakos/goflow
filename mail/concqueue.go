package mail

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type ConcQueue struct {
	items list.List
	mu    sync.Mutex
	size  atomic.Uint32
}

func NewConcQueue(size int) *ConcQueue {
	s := atomic.Uint32{}
	s.Store(uint32(size))

	return &ConcQueue{
		items: *list.New(),
		size:  s,
	}
}

func (q *ConcQueue) Push(v int) {
	q.pus
}
