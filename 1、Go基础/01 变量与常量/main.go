package main

import "fmt"

// 1.变量与常量
// 常量
const pi = 3.1415

func main() {
	// 变量
	// 声明变量的三种方式
	var name1 string
	name1 = "tank"
	fmt.Println(name1)

	var name2 string = "tank jam"
	fmt.Println(name2)

	name3 := "tank jam jam "
	fmt.Println(name3)
	fmt.Println(pi)

	// 批量声明
	var (
		name   string
		age    int
		gender bool
	)
	name, age, gender = "tank", 18, false
	if gender {
		gender2 := "男"
		fmt.Printf("姓名: %s  年龄: %d   性别: %s\n", name, age, gender2)
	} else {
		gender2 := "女"
		fmt.Printf("姓名: %s  年龄: %d   性别: %s\n", name, age, gender2)
	}

}
