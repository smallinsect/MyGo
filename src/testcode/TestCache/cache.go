package testcache

import (
	"container/list"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"

	"m1-server/TestWorkPool"

	log "github.com/sirupsen/logrus"
)

type LruTable struct {
	rw        *sync.RWMutex
	size      int
	entryList *list.List //  实体链表
	items     *sync.Map  //  具体数据
	newEntryI func() EntryI
}

type EntryI interface {
	Save() (err error)
	SaveDataStatusUpdate()  // 更新保存数据到数据库状态
	SaveUpdateStatus() bool // 是否需要保存到数据库
	Load(key interface{}) (err error)
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewLruTable(size int, entryI func() EntryI) *LruTable {
	return &LruTable{size: size, entryList: list.New(), items: new(sync.Map), newEntryI: entryI}
}

func (lru *LruTable) Get(key interface{}) (val interface{}, ok bool) {
	if item0, ok := lru.items.Load(key); ok {
		item := item0.(*list.Element)
		val = item.Value.(*entry).value
		if val != nil {
			lru.entryList.MoveToFront(item)
			return val, true
		}
		lru.removeElement(item, true)
	}
	ent := lru.newEntryI()
	err := ent.Load(key)
	if err != nil {
		return
	}
	lru.Put(key, ent)
	return ent, true
}

func (lru *LruTable) GetForLoad(key interface{}) (val interface{}, isLoad, ok bool) {
	if item0, ok := lru.items.Load(key); ok {
		item := item0.(*list.Element)
		lru.entryList.MoveToFront(item)
		return item.Value.(*entry).value, false, true
	}
	ent := lru.newEntryI()
	err := ent.Load(key)
	if err != nil {
		return
	}
	lru.Put(key, ent)
	return ent, true, true
}

func (lru *LruTable) Put(key interface{}, value interface{}) {
	if item0, ok := lru.items.Load(key); ok {
		item := item0.(*list.Element)
		lru.entryList.MoveToFront(item)
		ent := item.Value.(*entry)
		ent.value = value
		return
	}
	element := lru.entryList.PushFront(&entry{key: key, value: value})
	lru.items.Store(key, element) // 保存
	if lru.size > 0 && lru.size < lru.entryList.Len() {
		lru.removeOldItem()
	}
}

func (lru *LruTable) Remove(key interface{}) {
	if item0, ok := lru.items.Load(key); ok && item0 != nil {
		item := item0.(*list.Element)
		if item != nil {
			lru.removeElement(item, false)
		}

	}
}

func (lru *LruTable) Save(w *workpool.WorkPool) {
	for item := lru.entryList.Front(); item != nil; item = item.Next() {
		if item.Value == nil {
			continue
		}
		if item.Value.(*entry).value == nil {
			continue
		}
		val := item.Value.(*entry).value.(EntryI)
		if val != nil {
			if val.SaveUpdateStatus() {
				f := func(v EntryI) {
					v.Save()
					v.SaveDataStatusUpdate()
				}
				f1 := func() {
					defer func() {
						if err := recover(); err != nil {
							log.Error("stacktrace from panic lrutable save func: \n", string(debug.Stack()))
						}
					}()
					f(val)
				}
				if !w.Do(f1) {
					break
				}
			}
		}
	}
}

func (lru *LruTable) SaveAll() {
	ch := make(chan int, 50) // 并发50像数据库写入数据
	wg := new(sync.WaitGroup)
	for item := lru.entryList.Front(); item != nil; item = item.Next() {
		if item.Value == nil {
			continue
		}
		if item.Value.(*entry).value == nil {
			continue
		}
		val := item.Value.(*entry).value.(EntryI)
		if val != nil {
			wg.Add(1)
			ch <- 1
			go func(v EntryI) {
				defer func() {
					<-ch
					wg.Done()
				}()
				v.Save()
				v.SaveDataStatusUpdate()
			}(val)

		}
	}
	wg.Wait()
}

func (lru *LruTable) removeOldItem() {
	lastItem := lru.entryList.Back()
	if lastItem != nil {
		lru.removeElement(lastItem, true)
	}
}

func (lru *LruTable) removeElement(item *list.Element, saveDb bool) {
	lru.entryList.Remove(item)
	if item.Value == nil {
		return
	}
	ent := item.Value.(*entry)
	lru.items.Delete(ent.key)
	if saveDb {
		ent.value.(EntryI).Save()
	}

}

// Resize changes the cache size.
func (lru *LruTable) Resize(size int) (evicted int) {
	diff := lru.entryList.Len() - size
	if diff < 0 {
		diff = 0
	}
	for i := 0; i < diff && size > 0; i++ {
		lru.removeOldItem()
	}
	lru.size = size
	return diff
}

type Cache struct {
	data     *sync.Map
	workPool *workpool.WorkPool
}

func NewCache() *Cache {
	return &Cache{data: new(sync.Map), workPool: workpool.NewWorkPool(100, 500)}
}

func (c *Cache) Get(table string, key interface{}) (value interface{}, ok bool) {
	v, ok := c.data.Load(table)
	if !ok {
		return
	}
	return v.(*LruTable).Get(key)
}

// isLoad 是否从数据库加载了数据
func (c *Cache) GetForLoad(table string, key interface{}) (value interface{}, isLoad, ok bool) {
	v, ok := c.data.Load(table)
	if !ok {
		return
	}
	return v.(*LruTable).GetForLoad(key)
}

func (c *Cache) Put(table string, key interface{}, value interface{}) {
	v, ok := c.data.Load(table)
	if !ok {
		panic("not found " + table)
		return
	}
	v.(*LruTable).Put(key, value)
	return
}

func (c *Cache) Resize(table string, size int) (evicted int, err error) {
	v, ok := c.data.Load(table)
	if !ok {
		err = errors.New("cache 没有 " + table)
		return
	}
	evicted = v.(*LruTable).Resize(size)
	return
}

func (c *Cache) Remove(table string, key interface{}) (err error) {
	v, ok := c.data.Load(table)
	if !ok {
		err = errors.New("cache 没有 " + table)
		return
	}
	v.(*LruTable).Remove(key)
	return
}

func (c *Cache) AddTable(table string, size int, entryI func() EntryI) {
	c.data.Store(table, NewLruTable(size, entryI))
}

func (c *Cache) SaveTable(table string) (err error) {
	v, ok := c.data.Load(table)
	if !ok {
		err = errors.New("cache 没有 " + table)
		return
	}
	v.(*LruTable).Save(c.workPool)
	return
}

func (c *Cache) Save() {
	c.data.Range(func(key interface{}, val interface{}) bool {
		val1 := val.(*LruTable)

		val1.Save(c.workPool)
		return true
	})

}

func (c *Cache) SaveAll() {
	fmt.Println("save--->all 22222.")
	c.data.Range(func(key interface{}, val interface{}) bool {
		val1 := val.(*LruTable)
		val1.SaveAll()
		return true
	})
}
