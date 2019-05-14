package main

import "fmt"

func main() {
	//定义输入项
	key := ""

	//定义软件退出条件
	loop := true
	//定义家庭账户余额
	balance := 10000.0
	//定义每笔收支金额
	money := 0.0
	//定义收支明细
	var note string
	//收支详情。当有收支时对details进行拼接
	details := "收支\t账户余额\t收支金额\t说    明\n"
	//是否有收支行为的flag
	flag := false

	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&key)

		switch key {
		case "1":
			if flag {
				fmt.Println("-----------------当前收支明细-----------------")
				fmt.Println(details)
			} else {
				fmt.Println("还没有收支情况，来一笔吧？1、支出；2、收入")
			}
		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balance += money //修改账户余额
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("收入\t%v\t%v\t%v\n", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("你的余额不足，请重新输入！！！")
				break
			}
			balance -= money
			fmt.Println("本次支出明细")
			fmt.Scanln(&note)
			details += fmt.Sprintf("支出\t%v\t%v\t%v\n", balance, money, note)
			flag = true
		case "4":
			fmt.Println("你确定要退出本软件吗？y/n")
			for {
				loop2 := false
				choise := ""
				fmt.Scanln(&choise)
				switch choise {
				case "y":
					loop = false
					loop2 = true
				case "n":
					loop = true
					loop2 = true
				default:
					fmt.Println("你的输入错误，请重新输入。y/n")
				}
				if loop2 {
					break
				}
			}
		default:
			fmt.Println("请输入正确的选项")
		}

		if !loop {
			break
		}
	}
	fmt.Println("你已经退出了家庭收支软件....")
}
