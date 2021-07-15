package main

import "fmt"

// 自定义类型和类型的别名提示

// 1.自定义类型
// MyInt
type MyInt int

// 2.类型别名
type NewInt = int

func main() {
	var a MyInt
	var b NewInt
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
