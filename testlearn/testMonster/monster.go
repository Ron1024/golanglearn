package testmonster

import (
	"encoding/json"
	"fmt"
	"os"
)

//Monster 怪物类
type Monster struct {
	Name   string            `json:"name"`
	Age    int               `json:"age"`
	Addr   map[string]string `json:"addr"`
	Skills []string          `json:"skills"`
}

//NewMonster 初始化Monster
func NewMonster(name string, age int, addr map[string]string, skills []string) *Monster {
	return &Monster{
		Name:   name,
		Age:    age,
		Addr:   addr,
		Skills: skills,
	}
}

//MonsterServizial 怪物序列化
func (monster *Monster) MonsterServizial() {
	filename := ""
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()
	str, err := json.Marshal(monster)
}

//MonsterUnServizial 怪物反序列化
func (monster *Monster) MonsterUnServizial() {

}
