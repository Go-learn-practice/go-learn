package basic

import "fmt"

func VarDeclareCase() {
	// 通过var关键词声明变量
	var i int
	var j = 100

	fmt.Println(i, j)

	// 数组
	var arr = [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{1, 2, 3, 4, 5}
	var arr2 [5]int
	arr2[2] = 4
	arr2[3] = 5
	fmt.Println(arr, arr1, arr2)

	// 指针类型，用于表示变量地址的类型
	var intPtr *int
	var floatPtr *float64

	var i1 = 100
	fmt.Println(f1(&i1))
	fmt.Println(i1)

	fmt.Println(intPtr, floatPtr)
}

func f1(i *int) int {
	*i = *i + 1

	return *i
}
