package main

import "fmt"

/*
	channel 初体验
	- 定义: c1 := make(chan int, 10)
	- 基本操作:
		- 发送
			- c1 <- 10
		- 接收
			- ret := <- c1
		- 关闭
			- close(c1)
*/

func main() {
	// 定义一个ch1变量
	var ch1 chan int  // channel内部职能传int类型
	var ch2 chan string
	fmt.Println(ch1)
	fmt.Println(ch2)

	// channel是引用类型，所以需要通过make来初始化s
	ch3 := make(chan int, 10)  // 初始化内存分配

	// 通道的操作: 发送、接收、关闭
	ch3 <- 10
	ret := <- ch3
	fmt.Println(ret)

	ch3 <- 20
	ch3 <- 30

	close(ch3)

	// 关闭通道的特点
	// 1.若 channel 关闭后再取值，则会取类型对应的默认值
	ret2 := <- ch3  // close
	fmt.Println(ret2)

	// 2.往关闭的通道中插入数据，会报错
	// panic: send on closed channel
	//ch3 <- 100

	// 3.关闭 已关闭的通道会报错
	// panic: close of closed channel
	//close(ch3)

}
