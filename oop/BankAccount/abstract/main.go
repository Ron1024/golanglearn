package main

import "fmt"

//Account 银行账户信息
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//Query 查询账号余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你的密码不正确")
	}
	fmt.Printf("你的账号为：%v	余额为：%v\n", account.AccountNo, account.Balance)
}

//GetMoney 取款
//1、验证密码
//2、查询余额
//3、扣除余额，取款成功
func (account *Account) GetMoney(pwd string, money float64) {
	if pwd != account.Pwd {
		fmt.Println("你的密码不正确")
		return
	}
	if money > account.Balance {
		fmt.Println("你输入的金额不能取出")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

func main() {
	//验证取款流程
	//1、初始化存款账户
	account := &Account{
		AccountNo: "123",
		Pwd:       "456",
		Balance:   99.9,
	}
	//2、验证密码
	account.Query("456")
	//3、验证取款
	account.GetMoney("456", 100.0)

}
