package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/parker/go/src/github.com/parker/golanglearn/FIleOpt/example01/filedemo02/hello.txt", os.O_WRONLY|os.O_APPEND, 06666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	file.WriteString("锤子就是锤子\n")
	writer.Flush()
}
