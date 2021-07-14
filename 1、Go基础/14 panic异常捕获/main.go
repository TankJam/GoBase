package main

import "fmt"

/*
	panic: 主动抛出异常，自定义输出错误信息
*/

func a(){
	fmt.Println("func b print...")
}

func b() {
	// defer + panic 在异常触发后执行某些操作
	defer func() {
		// recover: 捕获异常信息;
		err := recover()  // 捕获b函数中panic异常信息
		if err != nil{
			fmt.Println("func b errors")
			fmt.Println(err)
			fmt.Printf("%T\n", err)
		}
	}()
	// 主动触发异常，然后打印自定义异常信息，再去执行defer里面的代码
	panic("panic in b")
}

func main() {
	b()
}
