package service

import "github.com/parker/golanglearn/CustomerManage/moudle"

//CustomerService 数据服务类
type CustomerService struct {
	customers   []moudle.Customer
	customerNum int
}

//NewCustomerService 工厂类，向外提供初始化方法
func NewCustomerService() *CustomerService {
	cs := &CustomerService{}
	cs.customerNum = 1
	customer := moudle.NewCustomer(1, "张三", "男", 20, "112", "123@qq.com")
	cs.customers = append(cs.customers, customer)
	return cs
}

//List 返回用户切片
func (cv *CustomerService) List() []moudle.Customer {
	return cv.customers
}

//AddCustomer 添加客户
func (cv *CustomerService) AddCustomer(customer moudle.Customer) bool {
	customer.ID = cv.customerNum + 1
	cv.customers = append(cv.customers, customer)
	return true
}

//FindCustomerByID 根据ID在customer切片中查找客户信息，如果没有返回-1,如果有返回该客户
func (cv *CustomerService)FindCustomerByID(id int) int  {
	index :=-1
	for i:= 0 ;i<len(cv.customers);i++{
		if cv.customers[i].ID == id {
			index = i
		}
	}
	return index
}

//DeleteCustomerByID 根据ID在Customer切片中删除customer
//返回是否删除成功标签
func (cv *CustomerService) DeleteCustomerByID(id int) bool {
	//flag 删除成功标记
	flag := false
	for i := 0 ; i< len(cv.customers);i++{
		if cv.customers[i].ID==id {
			cv.customers = append(cv.customers[:i],cv.customers[i+1:]...)
			flag = true
			break
		}
	}
	return flag
}

//ModifyCustomerByID 根据ID修改customer信息
func (cv *CustomerService)  ModifyCustomerByID(customer moudle.Customer)bool {
	flag := false
	for i := 0; i< len(cv.customers);  i++{
		if cv.customers[i].ID ==customer.ID {
			cv.customers[i].Name = customer.Name
			cv.customers[i].Age = customer.Age
			cv.customers[i].Gender = customer.Gender
			cv.customers[i].Phone = customer.Phone
			cv.customers[i].Email = customer.Email
			flag = true
		}
	}
	return flag
}