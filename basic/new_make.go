package basic

import "fmt"

func NewCase() {
	// 通过new函数，可以创建任意类型，并返回指针
	mpPtr := new(map[string]*user)
	// panic: assignment to entry in nil map
	// (*mpPtr)["A"] = &user{}
	if *mpPtr == nil {
		fmt.Println("map 为空")
	}
	fmt.Println(mpPtr)

	slicePtr := new([]user)
	// panic: runtime error: index out of range [0] with length 0
	// fmt.Println((*slicePtr)[0])

	*slicePtr = append(*slicePtr, user{Name: "nick"})
	if *slicePtr == nil {
		fmt.Println("切片值为 nil", *slicePtr)
	}
	fmt.Println(slicePtr, *slicePtr)
}

func MapInit1() {
	mpPtr := new(map[string]*user)
	// 初始化map
	*mpPtr = make(map[string]*user)
	(*mpPtr)["A"] = &user{Name: "A"}
	fmt.Println(*mpPtr, mpPtr)
}

func MapInit2() {
	mp := make(map[string]*user) // 创建并初始化 map
	mpPtr := &mp                 // 获取 map 的指针
	(*mpPtr)["A"] = &user{
		Name: "Alice",
		Age:  25,
		Addr: Address{
			Province: "SomeProvince",
			City:     "SomeCity",
		},
	}
	fmt.Println(*mpPtr, mpPtr)
}

// make 仅用于 切片、集合、通道的初始化

func MakeCase() {
	// 初始化切片，并设置长度和容量
	slice := make([]int, 10, 20)
	slice[0] = 10
	// 初始化集合，并设置集合的大小
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	// 初始化通道，设置通道的方向和缓冲大小
	// 可读写
	ch := make(chan int, 10)
	// 可写
	ch1 := make(chan<- int, 10)
	// 可读
	ch2 := make(<-chan int)

	fmt.Println(slice, mp, ch, ch1, ch2)
}

func SliceAndMapCase() {
	var slice []int
	slice = []int{1, 2, 3, 4, 5, 6}

	slice1 := make([]int, 10)
	slice1[1] = 10

	fmt.Println(slice, slice1)

	// 切片的截取
	slice2 := make([]int, 5, 10)
	fmt.Println(len(slice2), cap(slice2))
	slice2[0] = 0
	slice2[1] = 1
	slice2[2] = 2
	slice2[3] = 3
	slice2[4] = 4
	// 切片截取
	slice3 := slice2[2:5]
	fmt.Println(len(slice3), cap(slice3), slice3)

	slice4 := slice2[1:10]
	fmt.Println(len(slice4), cap(slice4), slice4)

	// 切片的附加
	slice4 = append(slice4, 1, 2, 3, 4, 5, 6)
	fmt.Println(len(slice4), cap(slice4), slice4)

	// 集合，无序
	mp := make(map[string]string, 10)
	mp["A"] = "a"
	mp["B"] = "b"
	mp["C"] = "c"
	mp["D"] = "d"
	fmt.Println(mp)
	for k, v := range mp {
		fmt.Println(k, v)
	}
	// 删除集合元素
	delete(mp, "B")
	fmt.Println(mp)
}
