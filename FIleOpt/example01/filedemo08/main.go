package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("/usr")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(info)
}
