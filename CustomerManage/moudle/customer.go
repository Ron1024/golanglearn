package moudle

import "fmt"

//Customer 定义客户对象
type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//NewCustomer 初始化客户信息
func NewCustomer(id int, name, gender string, age int, phone, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//GetCustomer 格式化客户信息
func (customer *Customer) GetCustomer() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", customer.ID, customer.Name, customer.Gender, customer.Age, customer.Phone, customer.Email)
	return info
}
