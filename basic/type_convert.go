package basic

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func ConvertCase() {
	// 同一类型转换 数字和数字、字符串和字符和字节
	// 不同类型的数据转换 数字和字符串
	// 接口类型转其他类型

	// 数字类型的转换
	var num1 int = 100
	fmt.Println(int64(num1))

	// 字符串与数字类型的转换
	var num3 = 100
	fmt.Println(strconv.Itoa(num3) + "abc")
	var str1 = "100"
	fmt.Println(strconv.Atoi(str1))

	var num4 int64 = 1010
	fmt.Println(strconv.FormatInt(num4, 10))
	var str2 = "1010"
	fmt.Println(strconv.ParseInt(str2, 10, 64))

	// 字符串与[]byte转换
	var str3 = "今天天气很好"
	byte1 := []byte(str3)
	fmt.Println(byte1)

	fmt.Println(string(byte1))

	// 字符串与rune转换
	var rune1 = []rune(str3)
	fmt.Println(rune1)
	fmt.Println(string(rune1))

	// 接口类型转具体类型
	var inf interface{} = 100
	var infStruct interface{} = user{}
	i, ok := inf.(int)
	fmt.Println(ok, i)
	u, ok := infStruct.(user)
	fmt.Println(ok, u)

	// 时间转字符串
	var t time.Time
	t = time.Now()
	timeStr := t.Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
	// 字符串转时间
	t2, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	fmt.Println(t2)

	// uintptr
	u1 := user{}
	// unsafe.Pointer 是一个通用的指针类型，不能用于计算
	uPtr := unsafe.Pointer(&u1)
	namePtr := unsafe.Pointer(uintptr(uPtr) + unsafe.Offsetof(u1.Name))
	*(*string)(namePtr) = "nick"
	fmt.Println(u1)
}
