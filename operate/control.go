package operate

import "fmt"

func FlowControlCase() {
	ifElseCase(0)
	ifElseCase(1)
	ifElseCase(2)
	forCase()
	switchCase("C", 1)
}

func ifElseCase(a int) {
	if a == 0 {
		fmt.Println("执行 if 语句块")
	} else if a == 1 {
		fmt.Println("执行 else if 语句块")
	} else {
		fmt.Println("执行 else 语句块")
	}
}

func forCase() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("位置 1 执行 for 语句块 i：", i)
	}

	var i int
	var b = true
	for b {
		fmt.Println("位置 2 执行 for 语句块 i：", i)
		i++
		if i >= 10 {
			b = false
		}
	}

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, value := range list {
		fmt.Printf("位置 3 执行 for 语句块 index：%d, data：%d \n", index, value)
	}

	mp := map[string]string{"A": "a", "B": "b", "C": "c"}
	for k, v := range mp {
		fmt.Printf("位置 4 执行 for 语句块 k：%s, v：%s \n", k, v)
	}

	str := "今天天气很好"
	for index, char := range str {
		fmt.Printf("位置 4 执行 for 语句块 index：%v, char：%v \n", index, string(char))
	}
}

func switchCase(alpha string, in interface{}) {
	switch alpha {
	case "A":
		fmt.Println("执行 A 语句块")
	case "B":
		fmt.Println("执行 B 语句块")
	case "C", "D":
		fmt.Println("执行 CD 语句块")
		fallthrough
	case "E":
		fmt.Println("执行 E 语句块")
	case "F":
		fmt.Println("执行 F 语句块")
	}

	switch in.(type) {
	case string:
		fmt.Println("in 输入为字符串")
	case int:
		fmt.Println("in 输入为 int 类型")
	default:
		fmt.Println("未能识别 in 类型")
	}
}
