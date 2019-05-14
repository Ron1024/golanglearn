package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil。Readfile一次性将文件读取到位
	file := "/Users/parker/go/src/github.com/parker/golanglearn/FIleOpt/example01/filedemo02/hello.txt"
	countent, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("读取文件出错了：", err)
	}
	fmt.Print("%v", countent) //[]byte
	fmt.Print("%v", string(countent))
}
