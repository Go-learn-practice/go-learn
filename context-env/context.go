package context_env

import (
	"context"
	"fmt"
	"time"
)

func ContextSelect() {
	done := make(chan struct{})
	go f1(done)
	go f1(done)
	time.Sleep(time.Second * 3)
	close(done)
}

func f1(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("协程退出")
			return
		}
	}
}

func ContextCase() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "desc", "ContextCase")
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	data := [][]int{{1, 2}, {3, 2}}
	ch := make(chan []int)
	go calculate(ctx, ch)

	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}
	time.Sleep(time.Second * 10)
}

func calculate(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			ctx := context.WithValue(ctx, "desc", "calculate")
			ch := make(chan []int)
			go sumContext(ctx, ch)
			ch <- item

			ch1 := make(chan []int)
			go multiContext(ctx, ch1)
			ch1 <- item
			fmt.Println(item)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("calculate 协程退出 context desc: %s, 错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func sumContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := sum(a, b)
			fmt.Printf("%d + %d = %d\n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("sumContext 协程退出 context desc: %s, 错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func multiContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := multi(a, b)
			fmt.Printf("%d * %d = %d\n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("multiContext 协程退出 context desc: %s, 错误消息：%s\n", desc, ctx.Err())
			return
		}
	}
}

func sum(a, b int) int {
	return a + b
}

func multi(a, b int) int {
	return a * b
}
