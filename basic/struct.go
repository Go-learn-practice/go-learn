package basic

import "fmt"

type user struct {
	Name string
	Age  uint
	Addr Address
}

type Address struct {
	Province string
	City     string
}

func StructCase() {
	// 值类型
	u := user{
		Name: "nick",
		Age:  18,
	}
	f2(u)
	fmt.Println(u)

	// 指针类型
	m := &user{
		Name: "jk",
		Age:  20,
	}
	f3(m)
	fmt.Println(m)

	// 指针类型
	m1 := new(user)
	m1.Name = "tony"
	m1.Age = 21
	fmt.Println(m1)

	// 结构体为值类型，定义变量后被默认触发
	var u1 user
	fmt.Println(u1)
}

func f2(u user) {
	u.Age = 28
}

func f3(u *user) {
	u.Age = 24
}
