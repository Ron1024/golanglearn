package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	//1、file 叫file对象
	//2、file 叫file指针
	//3、file 叫file文件句柄
	file, err := os.Open("/Users/parker/go/src/github.com/parker/golanglearn/familyAccountsingle/testAccount.go")
	if err != nil {
		fmt.Println("打开文件失败:", err)
	}
	fmt.Printf("%v", file)

	errs := file.Close()
	if err != nil {
		fmt.Println("file close fail：", errs)
	}
}
