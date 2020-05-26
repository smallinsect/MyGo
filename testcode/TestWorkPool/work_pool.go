package workpool

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type TaskFunc func()

type WorkPool struct {
	closed    int32
	isQueTask bool // 标记是否队列取出任务
	wg        sync.WaitGroup
	task      chan TaskFunc
	waitQueue *SyncQueue // 等待队列
}

var defaultWorkPool *WorkPool

// 不要把 带有阻塞的任务放到这里
func Do(f TaskFunc) {
	if defaultWorkPool == nil {
		defaultWorkPool = NewWorkPool(1000, 100000)
	}
	defaultWorkPool.Do(f)
}

func WaitClose() {
	if defaultWorkPool == nil {
		return
	}
	defaultWorkPool.WaitClose()
}

func NewWorkPool(max int, queueSize int) *WorkPool {
	if max < 1 {
		max = 1
	}
	p := &WorkPool{
		task:      make(chan TaskFunc, 2*max),
		waitQueue: NewSyncQueue(queueSize),
	}
	go p.run(max)
	return p
}

func (p *WorkPool) Do(fn TaskFunc) bool { // 添加到工作池，并立即返回
	if p.IsClosed() { // 已关闭
		return false
	}
	return p.waitQueue.Push(fn)
}

// 添加到工作池，并等待执行完成之后再返回
func (p *WorkPool) DoWait(task TaskFunc) {
	if p.IsClosed() { // closed
		return
	}

	doneChan := make(chan struct{})
	p.waitQueue.Push(TaskFunc(func() {
		defer close(doneChan)
		task()
	}))
	<-doneChan
}

func (p *WorkPool) WaitClose() { // 等待工作线程执行结束
	p.waitQueue.Wait() // 等待队列结束
	p.waitQueTask()    // 等待队列中的任务都结束
	close(p.task)
	atomic.StoreInt32(&p.closed, 1)
	p.wg.Wait() // 等待结束
}

// 是否已经关闭
func (p *WorkPool) IsClosed() bool {
	if atomic.LoadInt32(&p.closed) == 1 {
		return true
	}
	return false
}

func (p *WorkPool) startQueue() {
	for {
		fn := p.waitQueue.Pop().(TaskFunc)
		p.isQueTask = true
		if p.IsClosed() { // closed
			p.waitQueue.Close()
			p.isQueTask = false
			break
		}

		if fn != nil {
			p.task <- fn
		}
		p.isQueTask = false
	}
}

func (p *WorkPool) waitQueTask() {
	for {
		runtime.Gosched() // 出让时间片
		if !p.isQueTask {
			break
		}
	}
}

func (p *WorkPool) run(maxWorkerNum int) {
	go p.startQueue() // 启动队列

	p.wg.Add(maxWorkerNum) // 最大的工作协程数
	// 启动max个worker
	for i := 0; i < maxWorkerNum; i++ {
		go func() {
			defer p.wg.Done()
			// worker 开始干活
			for wt := range p.task {
				if wt == nil {
					continue // 需要先消费完了之后再返回，
				}
				wt() // 真正运行
			}
		}()
	}
}
