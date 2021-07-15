package main

import "fmt"

/*
结构体: 类似于其他语言的类;
	- 定义语法:
		type 类型名 struct{
			字段名1 字段类型
			字段名2 字段类型
		}

	- 类型名: 自定义结构体的名称，在同一个包内不可重复;
	- 字段名: 结构体字段名，结构体中字段名是唯一的;
	- 字段类型: 结构体字段的具体类型;

匿名结构体: 一次性定义与使用的结构体;
*/


// 结构体定义
type person struct{
	name, sex string
	age int8
}

func main() {
	// 使用结构体得到空对象
	var pObj person
	// 给对象的属性添加值
	pObj.name = "tank"
	pObj.age = 18
	pObj.sex = "male"
	fmt.Println(pObj, pObj.name, pObj.age, pObj.sex)

	// 匿名结构体的定义与使用
	var user struct{
		name string
		married bool
	}
	user.name = "贝塔"
	user.married = false
	fmt.Println(user, user.name, user.married)
}
