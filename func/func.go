package _func

import "errors"

func SumCase(a, b int) (sum int, err error) {
	if a <= 0 && b <= 0 {
		err = errors.New("两数相加benign同时小于0")
	}
	sum = a + b
	return
}

func ReferenceCase(a int, b *int) {
	a += 1
	*b += 1
}

var g int
var G int

func ScopeCase(a, b int) {
	c := 100
	g = a + b + c
	G = g
}

type User struct {
	Name string
	Age  uint
}

func NewUser(name string, age uint) *User {
	return &User{name, age}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() uint {
	return u.Age
}
