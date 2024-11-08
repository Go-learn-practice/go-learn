package standard

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func EncodingCase() {
	type user struct {
		ID   int64
		Name string
		Age  uint8
	}

	u := user{ID: 1, Name: "nick", Age: 18}

	fmt.Println("====json序列化====")
	// 序列化
	bytes, err := json.Marshal(u)
	fmt.Println(bytes, err)

	u1 := user{}
	// 反序列化
	err = json.Unmarshal(bytes, &u1)
	fmt.Println(u1, err)

	fmt.Println("====base64编码====")
	// base64编码
	str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println(str)
	bytes1, err := base64.StdEncoding.DecodeString(str)
	fmt.Println(bytes1, err)

	fmt.Println("====16进制编码====")
	// 16进制编码
	str1 := hex.EncodeToString(bytes1)
	fmt.Println(str1)
	bytes2, err := hex.DecodeString(str1)
	fmt.Println(bytes2, err)
}
