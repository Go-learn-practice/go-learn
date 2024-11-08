package generic

import "fmt"

type user struct {
	ID   int64
	Name string
	Age  uint8
}

type address struct {
	ID       int
	Province string
	City     string
}

/*
集合转列表
*/
func mapToList[K comparable, T any](mp map[K]T) []T {
	list := make([]T, len(mp))
	var i int
	for _, v := range mp {
		list[i] = v
		i++
	}
	return list
}

func myPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

func TTypeCase() {
	userMp := make(map[int64]user)
	userMp[1] = user{ID: 1, Name: "nick", Age: 18}
	userMp[2] = user{ID: 2, Name: "king", Age: 19}
	userList := mapToList[int64, user](userMp)

	ch := make(chan user)
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	addrMp := make(map[int64]address)
	addrMp[1] = address{ID: 1, Province: "广东", City: "深圳"}
	addrMp[2] = address{ID: 2, Province: "广东", City: "广州"}
	addrList := mapToList[int64, address](addrMp)

	ch1 := make(chan address)
	go myPrintln(ch1)
	for _, addr := range addrList {
		ch1 <- addr
	}
}

// List 泛型切片
type List[T any] []T

// MapT 泛型集合的定义
type MapT[K comparable, V any] map[K]V

// Chan 泛型通道的定义
type Chan[T any] chan T

func TTypeCase1() {
	userMp := make(MapT[int64, user])
	userMp[1] = user{ID: 1, Name: "nick", Age: 18}
	userMp[2] = user{ID: 2, Name: "king", Age: 19}
	var userList List[user]
	userList = mapToList[int64, user](userMp)

	ch := make(Chan[user])
	go myPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	addrMp := make(MapT[int64, address])
	addrMp[1] = address{ID: 1, Province: "广东", City: "深圳"}
	addrMp[2] = address{ID: 2, Province: "广东", City: "广州"}
	var addrList List[address]
	addrList = mapToList[int64, address](addrMp)

	ch1 := make(Chan[address])
	go myPrintln(ch1)
	for _, addr := range addrList {
		ch1 <- addr
	}
}
