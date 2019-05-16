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

func testStruct() {

	var addr map[string]string
	addr = make(map[string]string)
	addr["中国"] = "成都"
	addr["美国"] = "拉斯维加斯"
	skill := []string{"牛魔拳", "还我漂漂拳"}
	var monstor = &Monstor{
		Name:  "牛魔王",
		Age:   10,
		Addr:  addr,
		Skill: skill,
	}

	str, err := json.Marshal(monstor)
	if err != nil {
		fmt.Println("转换成为Json失败：", err)
		return
	}
	fmt.Println(string(str))

}

func main() {
	testStruct()
}
