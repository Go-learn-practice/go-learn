package standard

import (
	"fmt"
	"os"
)

func FmtCase() {
	// 打印到标准输出
	fmt.Println("今天天气很好")
	// 格式化，并打印到标准输出
	fmt.Printf("%s天气很好\n", "今天")
	// 格式化
	str := fmt.Sprintf("%s天气很好", "今天、明天")
	// 输出到io.writer
	fmt.Fprint(os.Stderr, str)
}

func FmtCase1() {
	type simple struct {
		value int
	}
	a := simple{
		value: 10,
	}
	// 通用占位符
	fmt.Printf("默认格式的值：%v \n", a)
	fmt.Printf("包含字段名的值：%+v \n", a)
	fmt.Printf("go语法表示的值：%#v \n", a)
	fmt.Printf("go语法表示的类型：%T \n", a)

	// 整数占位符
	v1 := 10
	v2 := 20170

	fmt.Printf("二进制：%b \n", v1)
	fmt.Printf("Unicode 码点转字符：%c \n", v2)
	fmt.Printf("十进制：%d \n", v1)
	fmt.Printf("八进制：%o \n", v1)
	fmt.Printf("0o为前缀的八进制 %0o \n", v1)
	fmt.Printf("用单引号将字符的值包起来：%q \n", v2)
	fmt.Printf("十六进制：%x \n", v1)
	fmt.Printf("十六进制大写：%X \n", v1)
	fmt.Printf("Uncoide格式：%U \n", v2)

	// 宽度设置

	// 浮点数占位

	// 字符串占位

	// 指针占位符
	str1 := "今天是个好日子"
	bytes := []byte(str1)
	// 切片第0个元素的地址
	fmt.Printf("%p \n", bytes)

	mp := make(map[string]int, 0)
	fmt.Printf("%p \n", mp)

	var p = new(map[string]int)
	fmt.Printf("%p \n", p)
}
