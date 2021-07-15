package main

import "fmt"

/*
	defer: 延迟执行，函数中有多个defer会从上到下去注册，
    当函数最后代码执行结束后，会从下到上执行 defer后面的代码或函数

*/


func main() {

	// 从上到下执行函数内代码, 并注册defer
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")

	// 从下到上执行defer
	// 打印顺序 start...  ---> end... ---> 3 ---> 2 ---> 1
}
