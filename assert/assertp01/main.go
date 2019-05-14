package main

import "fmt"

//Student 定义Student类型
type Student struct {
}

//TypeJudje 对类型进行断言
func TypeJudje(item ...interface{}) {
	for index, val := range item {
		switch val.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, val)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, val)
		case Student:
			fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, val)
		case *Student:
			fmt.Printf("第%v个参数是 *Student 类型，值是%v\n", index, val)

		}
	}
}

func main() {
	n1 := true
	n2 := "你好"
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudje(n1, n2, stu1, stu2)
}
