package wait_group

import (
	"fmt"
	"sync"
)

func MutexCase() {
	//singleRoutine()
	//multipleRoutine()
	//multipleSafeRoutine()
	multipleSafeRoutineByRWMutex()
}

// 单协程操作
func singleRoutine() {
	mp := make(map[string]int)
	list := []string{"A", "B", "C", "D", "E"}
	for i := 0; i < 20; i++ {
		for _, l := range list {
			_, ok := mp[l]
			if !ok {
				mp[l] = 0
			}
			mp[l] += 1
		}
	}
	fmt.Println(mp)
}

// 多协程操作，非协程安全
func multipleRoutine() {
	mp := make(map[string]int)
	list := []string{"A", "B", "C", "D", "E"}

	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, l := range list {
				_, ok := mp[l]
				if !ok {
					mp[l] = 0
				}
				mp[l] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp)
}

// 互斥锁协程安全
func multipleSafeRoutine() {
	type safeMap struct {
		data map[string]int
		sync.Mutex
	}
	mp := safeMap{
		data:  make(map[string]int),
		Mutex: sync.Mutex{},
	}
	list := []string{"A", "B", "C", "D", "E"}

	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mp.Lock()
			defer mp.Unlock()
			for _, l := range list {
				_, ok := mp.data[l]
				if !ok {
					mp.data[l] = 0
				}
				mp.data[l] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp.data)
}

type cache struct {
	data map[string]string
	sync.RWMutex
}

func newCache() *cache {
	return &cache{
		data:    make(map[string]string),
		RWMutex: sync.RWMutex{},
	}
}

func (c *cache) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	val, ok := c.data[key]
	if ok {
		return val
	}
	return ""
}

func (c *cache) Set(key, value string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

// 读写锁
func multipleSafeRoutineByRWMutex() {
	c := newCache()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Set("name", "nick")
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(c.Get("name"))
		}()
	}
	wg.Wait()
}
