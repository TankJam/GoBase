package main

import "fmt"

/*
	指针:
		- &值 ---> 地址
		- *地址 ---> 值

	注意: 因为Go语言函数中传参是值拷贝，
    相当于赋值一份并开启一个新的内存地址，
	若想对传递的参数进行修改，则需要使用指针。
*/

//// 1）*与&的基本使用
//func main() {
//	// a的值在内存中
//	a := 10
//	// 1.通过指针指向10的内存地址
//	b := &a
//	fmt.Println(b)  // 0xc0000180b8
//
//	// 2.通过内存地址得到值
//	c := *b
//	fmt.Println(c)  // 10
//}

// 2）函数传递参数，使用指针对传递的参数进行修改
// 在一个函数中修改另一个函数中参数的值
func modify(x int){
	x = 100
}

// 在一个函数中修改另一个函数中参数的指针
func modify2(y *int){
	*y = 100
}

func main() {
	a := 1
	modify(a)
	// a的值未修改，因为传递的是参数值
	fmt.Println(a)

	// a传递过去的是指针，所以会对原来的a变量值进行修改
	modify2(&a)
	fmt.Println(a)

}




