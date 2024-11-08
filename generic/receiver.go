package generic

import "fmt"

// MyStruct 泛型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

func (myStruct MyStruct[T]) GetData() T {
	return myStruct.Data
}

func ReceiverCase() {
	data := 18
	myStruct := MyStruct[*int]{
		Name: "nick",
		Data: &data,
	}
	data1 := myStruct.GetData()
	fmt.Println(data1, *data1)
}
