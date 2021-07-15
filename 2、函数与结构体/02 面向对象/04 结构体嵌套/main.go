package main

import "fmt"

/*
	结构体嵌套: （类似于Python的组合）
		- 结构体字段冲突问题
*/

// 结构体定义
type Address struct {
	Province string // 省份
	City string // 城市
	UpdateTime string  // 更新时间
}

type Email struct{
	Addr string  // 邮箱地址
	UpdateTime string  // 更新时间
}


// Person结构体中嵌套Address与Email结构体
type Person struct{
	Name string
	Gender string
	Age int8
	Address
	Email
}


func main() {
	p := Person{
		Name:    "tank",
		Gender:  "male",
		Age:     18,
		Address: Address{
			Province:   "广东",
			City:       "湛江",
			UpdateTime: "2021-07-15",
		},
		Email:   Email{
			Addr:       "15622792660@163.com",
			UpdateTime: "2021-07-15",
		},
	}

	fmt.Println(p.Name, p.Gender, p.Age, p.Address.Province)
}
