package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/parker/go/src/github.com/parker/golanglearn/FIleOpt/example01/filedemo02/hello.txt",
		os.O_RDWR|os.O_APPEND, 06666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf(str)
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {

		file.WriteString("锤子你好\n")
	}

	writer.Flush()
}
