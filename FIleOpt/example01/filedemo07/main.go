package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("/Users/parker/go/src/github.com/parker/golanglearn/FIleOpt/example01/filedemo02/hello.txt")
	if err != nil {
		fmt.Println("文件读取出错:", err)
		return
	}

	err = ioutil.WriteFile("/Users/parker/go/src/github.com/parker/golanglearn/FIleOpt/example01/filedemo07/hello.txt", data, 0666)
	if err != nil {
		fmt.Println("文件写入出错:", err)
		return
	}
}
