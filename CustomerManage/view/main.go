package main

import (
	"fmt"
	"github.com/parker/golanglearn/CustomerManage/moudle"
	"github.com/parker/golanglearn/CustomerManage/service"
)

//CustomerView 前端显示类
type CustomerView struct {
	key             string //用户输入信息
	loop            bool   //退出标记
	customerService *service.CustomerService
}

//list 获取所有的客户信息
func (cv *CustomerView) list() {
	//获取所有的客户信息，所有的信息都在customerservice的切片中
	customers := cv.customerService.List()
	fmt.Println("-----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetCustomer())
	}
	fmt.Printf("---------------客户列表完成----------------\n\n")
}
//add 添加客户
func (cv *CustomerView) add() {
	var name string
	var gender string
	var age int
	var phone string
	var email string
	fmt.Println("-----------------添加客户-----------------")
	fmt.Printf("姓名:")
	fmt.Scanf("%s", &name)
	fmt.Printf("性别:")
	fmt.Scanf("%s", &gender)
	fmt.Printf("年龄:")
	fmt.Scanf("%d", &age)
	fmt.Printf("电话:")
	fmt.Scanf("%s", &phone)
	fmt.Printf("邮箱:")
	fmt.Scanf("%s", &email)
	customer := moudle.Customer{
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
	success := cv.customerService.AddCustomer(customer)
	if success {
		fmt.Printf("---------------添加客户完成---------------\n\n")
	}else {
		fmt.Printf("---------------添加客户失败---------------\n\n")
	}
}
//delete 删除客户
func (cv *CustomerView) delete(){
	id := 0
	cv.list()
	fmt.Println("-----------------删除客户-----------------")
	fmt.Print("请输入你想删除的客户编号：")
	fmt.Scanf("%d",&id)
	//查找是否有该用户信息
	idx := cv.customerService.FindCustomerByID(id)
	if idx == -1 {
		fmt.Println("没有改用户信息，请重新输入。")
	}
	flag :=cv.customerService.DeleteCustomerByID(id)
	if flag{
		fmt.Print("---------------删除客户完成----------------\n\n")
	}else {
		fmt.Print("---------------删除客户失败----------------\n\n")
	}
}
//modify 更新客户
func (cv *CustomerView) modify()  {
	id := 0
	name:= ""
	gender := ""
	age := 0
	phone:=""
	email:=""
	cv.list()
	fmt.Println("-----------------修改客户-----------------")
	fmt.Print("请输入你想修改的客户编号：")
	fmt.Scanf("%d",&id)
	fmt.Print("请输入你想修改的客户姓名：")
	fmt.Scanf("%s",&name)
	fmt.Print("请输入你想修改的客户性别：")
	fmt.Scanf("%s",&gender)
	fmt.Print("请输入你想修改的客户年龄：")
	fmt.Scanf("%d",&age)
	fmt.Print("请输入你想修改的客户电话：")
	fmt.Scanf("%s",&phone)
	fmt.Print("请输入你想修改的客户邮箱：")
	fmt.Scanf("%s",&email)
	customer := moudle.Customer{
		ID:id,
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
	flag := cv.customerService.ModifyCustomerByID(customer)
	if flag {
		fmt.Println("---------------修改客户完成----------------")
	}else {
		fmt.Println("---------------修改客户失败----------------")
	}
}

//mainMenu 主菜单
func (cv *CustomerView) mainMenu() {
	for {
		fmt.Println("-----------------客户管理系统-----------------")
		fmt.Println("                  1 添加客户")
		fmt.Println("                  2 修改客户")
		fmt.Println("                  3 删除客户")
		fmt.Println("                  4 客户列表")
		fmt.Println("                  5 退    出")
		fmt.Println("请选择1~5:")
		fmt.Scanln(&cv.key)

		switch cv.key {
		case "1":
			cv.add()
		case "2":
			cv.modify()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("输入错误，请重新输入")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("你已经退出系统")
}
//主函数
func main() {
	cv := &CustomerView{
		loop: true,
		key:  "",
	}
	//初始化CustomerService
	cv.customerService = service.NewCustomerService()

	cv.mainMenu()
}
