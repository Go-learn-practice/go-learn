package channel

import (
	"fmt"
	"time"
)

// Communication 协程间通信
func Communication() {
	ch := make(chan int)
	go communicationF1(ch)
	go communicationF2(ch)
}

// F1 接收一个只写通道
func communicationF1(ch chan<- int) {
	// 通过循环向通道写数据
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

// F2 接收一个只读通道
func communicationF2(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// ConcurrentSync 并发场景下，同步机制
func ConcurrentSync() {
	// 带缓冲的通道
	ch := make(chan int, 10)
	// 向 ch 写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	// 从 ch 读取消息并打印
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}

// NoticeAndMultiplexing 通知协程退出与多路复用
func NoticeAndMultiplexing() {
	ch := make(chan int)
	strCh := make(chan string)
	done := make(chan struct{})
	go noticeAndMultiplexingF1(ch)
	go noticeAndMultiplexingF2(strCh)
	go noticeAndMultiplexingF3(ch, strCh, done)

	time.Sleep(5 * time.Second)
	close(done)
}

func noticeAndMultiplexingF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func noticeAndMultiplexingF2(ch chan<- string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("数字：%d", i)
	}
}

func noticeAndMultiplexingF3(ch <-chan int, strCh <-chan string, done <-chan struct{}) {
	i := 0
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case str := <-strCh:
			fmt.Println(str)
		case <-done:
			fmt.Println("收到退出通知，退出当前协程")
			return
		}
		i++
		fmt.Println("累计执行次数：", i)
	}
}
