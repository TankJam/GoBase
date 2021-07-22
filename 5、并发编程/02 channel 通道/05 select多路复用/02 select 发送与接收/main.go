package main

import "fmt"

/*
	select 往通道发送数据，并异步从通道接收出值
*/

func main() {
	// 声明一个存放int类型，容量为1的通道
	var ch1 = make(chan int, 1)

	// for 循环 10次
	for i := 0; i < 10; i++ {
		select {
		case ch1 <- i:  // 尝试往ch1中发送数据
		case ret := <- ch1:  // 尝试从ch1中接收数据
			fmt.Println(ret)
		}
	}
}
