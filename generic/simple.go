package generic

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var a1, b1 MyInt32 = 3, 4
	var c, d float64 = 5, 6

	// 编译器推断
	fmt.Println("使用泛型，数字比较：", getMaxNum(a, b))
	fmt.Println("使用泛型，数字比较：", getMaxCusNum[MyInt32](a1, b1))
	fmt.Println("使用泛型，数字比较：", getMaxNum[float64](c, d))

	var e, f int64 = 7, 8
	var g, h MyInt64 = 7, 8
	fmt.Println("使用泛型，数字比较：", getMaxCusNum(e, f))
	fmt.Println("使用泛型，数字比较：", getMaxCusNum(g, h))
}

func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CusNumT interface {
	// ～ 表示支持类型的衍生类型
	// | 表示取并集
	// 多行之间取交集
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint16
}

// MyInt64 为 int64 的衍生类型，是具有基础类型 int64 的新类型，与 int64 是不同类型
type MyInt64 int64

// MyInt32 为 int32 的别名 与 int32 同一类型
type MyInt32 = int32

func getMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 内置类型

func BuiltInCase() {
	var a, b string = "abc", "efg"
	fmt.Println("内置 comparable 泛型约束", getBuiltInComparable(a, b))
}

func getBuiltInComparable[T comparable](a, b T) bool {
	// comparable类型 仅支持 == != 两个操作
	if a == b {
		return true
	}
	return false
}
