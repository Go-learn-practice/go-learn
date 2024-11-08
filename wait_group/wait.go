package wait_group

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroupCase() {
	var a, b = 1000, 10000
	start := time.Now()
	for i := 0; i < 1000000000; i++ {
		multi(a, b)
	}
	t := time.Since(start)
	fmt.Println(t)

	start = time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50000000; j++ {
				multi(a, b)
			}
		}()
	}
	wg.Wait()
	t = time.Since(start)
	fmt.Println(t)
}

func WaitGroupCase1() {
	ch := make(chan []int, 1000)
	start := time.Now()

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		i := 0
		for item := range ch {
			fmt.Println(multi(item[0], item[1]))
			i++
		}
		time.Sleep(3 * time.Second)
		fmt.Println("数据处理完成，数据条数")
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		wg2.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			defer wg2.Done()
			for j := 0; j < 500; j++ {
				ch <- []int{i, j}
			}
		}(wg)
	}
	wg.Wait()
	close(ch)
	wg2.Wait()

	t := time.Since(start)
	fmt.Println("处理时长：", t)
}

func multi(a, b int) int {
	return a * b
}
