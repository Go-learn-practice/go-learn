package _interface

import "fmt"

type Dove struct {
	animal
}

func NewDove() Animal {
	return &Dove{}
}

func (a *Dove) Each() {
	fmt.Println("鸽子不吃")
}

func (a *Dove) Drink() {
	fmt.Println("鸽子不喝")
}

func (a *Dove) Sleep() {
	fmt.Println("鸽子不睡")
}

func (a *Dove) Run() {
	fmt.Println("鸽子不跑")
}
