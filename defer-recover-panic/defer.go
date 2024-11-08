package defer_recover_panic

import (
	"fmt"
	"log"
	"os"
)

/*
DeferCase1 关键词用来声明一个延迟调用函数 该函数可以是匿名函数也可以是具名函数
defer 延迟函数的执行顺序为后进先出
*/
func DeferCase1() {
	fmt.Println("开始执行DeferCase1")

	defer func() {
		fmt.Println("调用了匿名函数1")
	}()
	defer f1()
	defer func() {
		fmt.Println("调用了匿名函数2")
	}()

	fmt.Println("DeferCase1执行结束")
}

// DeferCase2 参与预计算
func DeferCase2() {
	i := 1

	// 传参
	defer func(j int) {
		fmt.Println("defer j: ", j)
	}(i + 1)

	// 闭包
	defer func() {
		fmt.Println("defer i: ", i)
	}()

	i = 99
	fmt.Println("i 的值：", i)
}

// DeferCase3 返回值
// defer 函数执行在 return 之后
func DeferCase3() {
	i, j := f2()
	fmt.Printf("i: %d, j: %d, g: %d \n", i, *j, g)
}

func ExceptionCase() {
	defer func() {
		err := recover()
		// 捕获异常 恢复协程
		if err != nil {
			fmt.Println("异常处理 defer recover", err)
		}
	}()
	fmt.Println("开始执行ExceptionCase方法")
	panic("ExceptionCase抛出异常")
}

func FileReadCase() {
	file, err := os.Open("../README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}

func f1() {
	fmt.Println("调用了具名函数1")
}

var g = 100

func f2() (int, *int) {
	defer func() {
		g = 200
	}()
	fmt.Println("f2 g: ", g)
	return g, &g
}
