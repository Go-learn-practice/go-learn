package basic

import "fmt"

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

type Gender uint8

const (
	FEMALE Gender = iota
	MALE
	THIRD
)

func ConstAndEnumCase() {
	const (
		A = 10
		B = 20
	)
	size := MB
	var gender Gender = MALE
	fmt.Println(A, B)
	fmt.Println(size)
	fmt.Println(gender)
}
