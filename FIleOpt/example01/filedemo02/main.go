package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	//1、file 叫file对象
	//2、file 叫file指针
	//3、file 叫file文件句柄
	file, err := os.Open("/Users/parker/go/src/github.com/parker/golanglearn/FIleOpt/example01/filedemo02/hello.txt")
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}
	//在函数退出时，要及时关闭file
	defer file.Close() //要及时关闭file句柄，否则会有内存泄露

	//创建一个带缓冲的*Reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF { //io.EOF为文件的末尾
			break
		}
	}

	fmt.Println("文件读取结束！")
}
