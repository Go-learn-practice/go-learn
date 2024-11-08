package _interface

import "fmt"

type Dog struct {
	animal
}

func NewDog() Animal {
	return Dog{}
}

func (a Dog) Each() {
	fmt.Println("狗吃狗粮")
}

func (a Dog) Drink() {
	fmt.Println("狗喝矿泉水")
}

func (a Dog) Sleep() {
	fmt.Println("狗不准睡觉")
}

func (a Dog) Run() {
	fmt.Println("狗不准跑")
}
