package wait_group

import (
	"fmt"
	"sync"
)

func MapCase() {
	mp := sync.Map{}
	// 设置键值对
	mp.Store("name", "nick")
	mp.Store("email", "nick@voice.com")
	// 通过key获取value，如果不存在就返回nil, ok 则返回false
	fmt.Println(mp.Load("name"))
	fmt.Println(mp.Load("email"))

	// 通过key获取value，如果不存在则设置指定的value并返回
	// ok 为 true 表示 key 存在并返回，为false表示 key 不存在并设置返回
	fmt.Println(mp.LoadOrStore("hobby", "篮球"))
	fmt.Println(mp.LoadOrStore("hobby", "羽毛球"))

	// 根据key获取value后，删除该key
	// ok 为true表示key存在，为false表示key不存在
	fmt.Println(mp.LoadAndDelete("hobby"))
	fmt.Println(mp.LoadAndDelete("hobby"))

	// 遍历mp
}
