package main

import (
	"encoding/json"
	"fmt"
)

//Monstor 定义一个monstor结构体
type Monstor struct {
	Name  string
	Age   int
	Addr  map[string]string
	Skill []string
}

func testUnMarshal() {
	str := "{\"Name\":\"牛魔王\",\"Age\":10,\"Addr\":{\"中国\":\"成都\",\"美国\":\"拉斯维加斯\"},\"Skill\":[\"牛魔拳\",\"还我漂漂拳\"]}"
	var monstor Monstor
	error := json.Unmarshal([]byte(str), &monstor)
	if error != nil {
		fmt.Println("反序列化失败：", error)
	}
	fmt.Println(monstor)
}

func main() {
	testUnMarshal()
}
