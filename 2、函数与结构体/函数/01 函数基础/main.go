package main

import "fmt"

/*
	- 有参与无参函数
	- 函数参数
		- 固定参数
		- 可变参数
	- 函数返回值
*/

// 无参函数 定义
func sayHello(){
	fmt.Println("hello")
}

// 有参函数 定义
func showName(name string){
	fmt.Println(name)
}

// 函数参数
/// 1.接收多个参数，并返回一个值
func getManyReturnOne(a int, b int)int{
	sumValue := a + b
	return sumValue
}

/// 2.接 收参数 与 返回值 简写
// (result int) 可以使 "return == return result"
func getManyReturnOne2(a, b, c int)(result int){  // 初始化变量名result
	result = a + b + c
	// 自动返回result
	return // "return == return result"
}

/// 3.接收 可变参数， 定义时 接收的参数名..., 表示可以接收可变参数
func getManyReturnOne3(a... int)(result int){
	result = 0
	for _, v := range a{
		result += v
	}
	return
}

// 函数返回值
/// 返回多个返回值
func getManyReturnMany(a, b int)(sum, sub int){
	sum = a + b
	sub = a - b
	return
}


func main() {
	// 无参函数 调用
	sayHello()

	// 有参函数调用
	showName("tank")

	// 接收参数与返回值
	ret1 := getManyReturnOne(10, 20)
	fmt.Println(ret1)

	// 接收参数与返回值 简写
	ret2 := getManyReturnOne2(100, 200, 300)
	fmt.Println(ret2)

	// 接收 可变长参数
	ret3 := getManyReturnOne3(1000, 2000, 3000, 4000, 5000, 6000)
	fmt.Println(ret3)

	// 多个返回值
	ret4, ret5 := getManyReturnMany(20000, 1000)
	fmt.Println(ret4, ret5)


}
