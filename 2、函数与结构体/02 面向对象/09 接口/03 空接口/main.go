package main

import (
	"fmt"
	"time"
)

/*
	空接口: 空接口的好处就是可以让接收指定类型变为接收任意类型的参数；
*/

// showType 接收空接口参数
func showType(a interface{}){
	fmt.Printf("type: %T\n", a)
}



func main() {
	showType("show type demo...")

	var x interface{}
	x = 100
	fmt.Println(x)
	x = "tank jam"
	fmt.Println(x)
	x = false
	fmt.Println(x)
	x = struct{name string}{name:"花花"}  // 匿名结构体
	fmt.Println(x)
	showType(time.Second)

	// 定义一个值为空的接口map
	userInfo := make(map[string]interface{}, 100)
	userInfo["张三"] = 100
	userInfo["李四"] = true
	userInfo["王五"] = 99.99
	userInfo["六六大顺"] = time.Now()
	fmt.Println(userInfo)
}
