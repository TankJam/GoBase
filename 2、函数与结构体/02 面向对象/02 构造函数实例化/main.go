package main

import "fmt"

/*
	造构函数: 构造一个可以 实例化 结构体的函数;
		- 类似于Python的 __init__

	注意: 结构体是值类型，使用 构造函数 + 结构体指针 进行构造;
*/

// 结构体
type person struct{
	name, city string
	age int8
}


// 构造函数
func initPerson(name, city string, age int8) person{
	return person{
		name: name,
		city: city,
		age:  age,  // 注意: 最后需要加逗号
	}
}


func main() {
	pObj := initPerson("tank", "上海", int8(18))
	fmt.Println(pObj, pObj.name, pObj.city, pObj.age)
	fmt.Printf("type: %T  value: %#v", pObj, pObj)
}
