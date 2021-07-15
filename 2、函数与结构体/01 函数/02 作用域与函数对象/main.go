package main

import "fmt"

/*
	函数变量的作用域
*/

// 1.定义全局变量
var number = 10

// 引用全局变量
func testGlobal(){
	fmt.Println(number)
}

// 2.在函数内部定义局部变量
func testNormals()int{
	//number := 200
	return number
}

// 3.函数对象
func test(x, y int, op func(int, int)int) int {
	// 调用接收的函数对象，并返回int
	return op(x, y)
}

func add(x, y int)int{
	return x + y
}

func sub(x, y int)int{
	return x - y
}


func main() {

	// 测试作用域
	testGlobal()
	ret1 := testNormals()
	fmt.Println(ret1)

	// 测试函数对象
	ret2 := test(100, 200, add)  // 传递函数对象add
	fmt.Println(ret2)

	ret3 := test(300, 200, sub)
	fmt.Println(ret3)

}
