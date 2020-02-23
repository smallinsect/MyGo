package testserver

import (
	"runtime"
	"sync"

	"gopkg.in/eapache/queue.v1"
)

// Synchronous FIFO queue
type SyncQueue struct {
	lock      sync.Mutex
	popable   *sync.Cond
	buffer    *queue.Queue
	closed    bool
	queueSize int
}

// Create a new SyncQueue
func NewSyncQueue(size int) *SyncQueue {
	ch := &SyncQueue{
		buffer:    queue.New(),
		queueSize: size,
	}
	ch.popable = sync.NewCond(&ch.lock)
	return ch
}

// Pop an item from SyncQueue, will block if SyncQueue is empty
func (q *SyncQueue) Pop() (v interface{}) {
	c := q.popable
	buffer := q.buffer

	q.lock.Lock()
	for buffer.Length() == 0 && !q.closed {
		c.Wait()
	}

	if buffer.Length() > 0 {
		v = buffer.Peek()
		buffer.Remove()
	}

	q.lock.Unlock()
	return
}

// Try to pop an item from SyncQueue, will return immediately with bool=false if SyncQueue is empty
func (q *SyncQueue) TryPop() (v interface{}, ok bool) {
	buffer := q.buffer

	q.lock.Lock()

	if buffer.Length() > 0 {
		v = buffer.Peek()
		buffer.Remove()
		ok = true
	} else if q.closed {
		ok = true
	}

	q.lock.Unlock()
	return
}

// Push an item to SyncQueue. Always returns immediately without blocking
func (q *SyncQueue) Push(v interface{}) bool {
	if q.Len() > q.queueSize {
		return false
	}
	q.lock.Lock()
	if !q.closed {
		q.buffer.Add(v)
		q.popable.Signal()
	}
	q.lock.Unlock()
	return true
}

// Get the length of SyncQueue
func (q *SyncQueue) Len() (l int) {
	q.lock.Lock()
	l = q.buffer.Length()
	q.lock.Unlock()
	return
}

func (q *SyncQueue) Close() {
	q.lock.Lock()
	if !q.closed {
		q.closed = true
		q.popable.Signal()
	}
	q.lock.Unlock()
}

//Wait 等待队列消费完成
func (q *SyncQueue) Wait() {
	for {
		if q.closed || q.buffer.Length() == 0 {
			break
		}
		runtime.Gosched() //出让时间片
	}
}