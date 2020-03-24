package testdemo

import (
	"fmt"
	"time"
)

//PoolFunc ...
func PoolFunc() {
	go func() {
		// 查询1
		fmt.Println("查询1")
	}()
	go func() {
		// 查询2
		fmt.Println("查询2")
	}()
	// ...........
	go func() {
		// 查询n
		fmt.Println("查询n")
	}()
}

type Grpool struct {
	queue chan uint8
	wg    *sync.WaitGroup
}

func (gp *Grpool) Add(n int) {
	for i := 0; i < n; i++ {
		gp.queue <- uint8(1)
	}
	gp.wg.Add(n)
}

func (gp *Grpool) Done() {
	<-gp.queue
	gp.wg.Done()
}

func (gp *Grpool) Wait() {
	gp.wg.Wait()
}

//PoolFunc01 ...
func PoolFunc01() {
	pool.Add(3)

	go func() {
		// 查询1
		pool.Done()
	}()

	go func() {
		// 查询2
		pool.Done()
	}()

	go func() {
		// 查询3
		pool.Done()
	}()

	pool.Wait()
}
