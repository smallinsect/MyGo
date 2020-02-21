package testserver

import (
	"fmt"
	"sync"
	"testing"
)

type Tcache struct {
	Val interface{}
}

func (tc *Tcache) Load(key interface{}) (err error) {
	tc.Val = key
	return
}

func (tc *Tcache) Save() (err error) {
	fmt.Println("save --> ", tc.Val)
	return
}

func NewTcache() EntryI {
	return &Tcache{}
}

func TestCache(t *testing.T) {
	c := NewCache()
	c.AddTable("tcache", 5, NewTcache)

	wg := new(sync.WaitGroup)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, ok := c.Get("tcache", "123")
			if !ok {
				//	fmt.Println("ok -->", ok)
				return
			}

			//fmt.Println("1->", v.(*Tcache).Val)
			v.(*Tcache).Val = "sb123"

			//fmt.Println("2->", v.(*Tcache).Val)
			c.Get("tcache", "456")
			c.Get("tcache", "asd")
			c.Get("tcache", "1234")
			v, ok = c.Get("tcache", "123")
			if !ok {
				//	fmt.Println("ok -->", ok)
				return
			}
			c.Put("tcache", "iii", &Tcache{1256})
			c.Put("tcache", "777", &Tcache{"8888"})
		}()
	}
	wg.Wait()

}
