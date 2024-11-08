package _interface

import "fmt"

type Cat struct {
	animal
}

func NewCat() Animal {
	return &Cat{}
}

func (c *Cat) Each() {
	fmt.Println("猫吃老鼠")
}
