package main

import (
	"fmt"
	"strings"
)

/*
	匿名函数: 定义一个无名字的函数，只能在定义时立马使用。
	闭包函数: 函数嵌套 + 函数对象 + 内层函数对外层函数变量的引用。
*/

// 匿名 + 闭包
func a(name string) func(){  // 返回值是一个函数对象
	// 返回一个匿名函数对象
	return func() {
		fmt.Println("hello tank jam...")
		fmt.Println(name)
	}
}

// 闭包实现：判断传入的名字是否以 .txt 作为后缀
func makeSuffixFunc(suffix string)func(string)string{
	return func(name string) string {
		// 若name的后缀不是suffix，就拼接成suffix后缀
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}else{
			return name
		}
	}
}

func main() {
	// 闭包函数: 内层函数引用外层函数变量
	ret1 := a("tank")
	fmt.Println(ret1)  // 返回的是函数地址
	ret1()

	// 闭包函数检测文件后缀
	ret2 := makeSuffixFunc(".txt")
	//ret3 := ret2("GoFiles.txt")
	ret3 := ret2("GoFiles")
	fmt.Println(ret3)


}
