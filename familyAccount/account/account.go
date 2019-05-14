package account

import "fmt"

//Account 账户信息
type Account struct {
	Balance float64
	Type    string
	Money   float64
	Note    string
	Details string
	Flag    bool
	Loop    bool
	Key     string
}

//QueryAccount 查询账户余额方法，与Account信息绑定
func (account *Account) QueryAccount() {
	if account.Flag {
		fmt.Println("-----------------当前收支明细-----------------")
		fmt.Println(account.Details)
	} else {
		fmt.Println("还没有收支情况，来一笔吧？")
	}
}

//InComingMoney 账户收入方法，与Account信息绑定
func (account *Account) InComingMoney() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&account.Money)
	account.Balance += account.Money //修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&account.Note)
	account.Details += fmt.Sprintf("收入\t%v\t%v\t%v\n", account.Details, account.Money, account.Note)
	account.Flag = true
}

//OutMoney 账户支出信息，与Account 信息绑定
func (account *Account) OutMoney() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&account.Money)
	if account.Money > account.Balance {
		fmt.Println("你的余额不足，请重新输入！！！")
	}
	account.Balance -= account.Money
	fmt.Println("本次支出明细")
	fmt.Scanln(&account.Note)
	account.Details += fmt.Sprintf("支出\t%v\t%v\t%v\n", account.Details, account.Money, account.Note)
	account.Flag = true
}

//Exit 退出软件方法
func (account *Account) Exit() {
	fmt.Println("你确定要退出本软件吗？y/n")
	for {
		loop2 := false
		choise := ""
		fmt.Scanln(&choise)
		switch choise {
		case "y":
			account.Loop = false
			loop2 = true
		case "n":
			account.Loop = false
			loop2 = true
		default:
			fmt.Println("你的输入错误，请重新输入。y/n")
		}
		if loop2 {
			break
		}
	}
}

//MainMenu 主菜单
func (account *Account) MainMenu() {
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&account.Key)

		switch account.Key {
		case "1":
			account.QueryAccount()
		case "2":
			account.InComingMoney()
		case "3":
			account.OutMoney()
		case "4":
			account.Exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !account.Loop {
			break
		}
	}
}

func NewAccount() *Account {
	return &Account{
		Key:     "",
		Loop:    true,
		Balance: 10000.0,
		Money:   0.0,
		Note:    "",
		Details: "收支\t账户余额\t收支金额\t说    明\n",
		Flag:    false,
	}
}
