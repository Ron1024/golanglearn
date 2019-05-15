package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

//创建一个新的文件，循环写入一些语句
	file ,err := os.OpenFile("/Users/parker/go/src/github.com/parker/golanglearn/FIleOpt/example01/filedemo04/hello.txt",os.O_WRONLY|os.O_CREATE,0666);
	if err!=nil {
		fmt.Println("打开文件错误:",err)
		return
	}
	defer file.Close()//及时关闭文件，防止内存泄露
	//准备写入5句话
	str := "Hellow Parker\n"
	//写入时使用带缓存的 *writer
	 writer :=bufio.NewWriter(file)
	for i:=0;i<5 ;i++  {
		writer.WriteString(str)
	}
	//因为 *writer是带缓存的，因此需要使用flush写入到磁盘
	writer.Flush()

} 
