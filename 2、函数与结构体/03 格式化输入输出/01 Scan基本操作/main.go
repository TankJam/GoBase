package main

import "fmt"

/*
	从终端获取用户输入的内容
	fmt.Scan(变量名1, 变量名2, 变量名3): 一次性输入多个值
	fmt.Scanln(变量名1): 一次性输入一个值
	fmt.Scanf(变量名1):按照固定的格式输入

注意: fmt.Scan 与 fmt.Scanf 不能一同使用。
*/
func main() {
	var (
		name string
		age int
		married bool
	)

	// 给 变量 传入用户输入的数据
	//fmt.Scan(&name, &age, &married)  // 一次性扫描多个变量
	//fmt.Printf("name:%s age:%02d married:%t\n", name, age, married)

	// "name:%s age:%02d married:%t\n"，用户输入的时候必须要根据format的格式去输入内容
	//fmt.Scanf("name:%s age:%02d married:%t\n", &name, &age, &married)

	fmt.Println(name, age, married)

	fmt.Println("请输入名字: ")
	fmt.Scanln(&name)  // 一次只能扫描一个变量

	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)  // 一次只能扫描一个变量

	fmt.Println("是否结婚: ")
	fmt.Scanln(&married)  // 一次只能扫描一个变量

	fmt.Println(name, age, married)

}

