package main

import "fmt"

/*
	new与make的区别:
		1.二者都是用来做内存分配的。
		2.make只用于slice、map、channel的初始化，返回的还是这三个引用类型本身。
		3.new用于值类型的内存分配，如int、string等,
          并且内存对应的值为类型零值，返回的是指向类型的指针。
*/

func main() {
	var a *int
	// int是值传递类型，指针类型 定义时是空指针
	// 注意: 不能对空指针进行赋值，需要先使用new来做初始化
	//*a = 100  // panic: runtime error: invalid memory address or nil pointer dereference
	a = new(int)
	*a = 100  //
	fmt.Println(*a)  // 100

	// map是引用传递类型，通过make定义时会开辟内存地址，可以修改属性
	var b = make(map[string]string)
	b["name"] = "tank jam"
	fmt.Println(b)
}
