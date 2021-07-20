package main

import (
	"fmt"
)

/*
	接收值时判断通道是否关闭
*/

func send(ch chan int){
	for i := 0; i < 10; i++ {
		ch <- i  // 往 channel 通道添加值
	}
	close(ch)  // channel被垃圾回收机制回收
}

func main() {
	var ch1 = make(chan int, 100)

	go send(ch1)  // 异步往管道发送请求

	// 1.利用for循环取通道ch1中接收值
	//for {
	//	// 使用 value, ok := <-ch1 取值方式，当通道关闭的时候 ok = false
	//	// 在send异步往通道添加值的时候，此处可以继续取值
	//	ret, ok := <-ch1
	//	if !ok {
	//		fmt.Println("channel is close...")
	//		break
	//	}
	//	fmt.Println(ret)
	//}

	// 2.利用for range 循环去ch1通道接收值
	for ret := range ch1{
		fmt.Println(ret)
	}
}
