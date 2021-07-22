package main

import "fmt"

/*
	- 无缓冲的通道：4*100交接棒，又称为同步通道
	- 有缓冲的通道：可以让程序实现异步操作
*/

func recv(ch chan bool){
	ret := <- ch  // 阻塞
	fmt.Println(ret)
}


func main() {

	ch := make(chan bool, 1)
	ch <- false
	// len: 获取数据量  cap: 获取容量
	fmt.Println(len(ch), cap(ch))

	// 调用recv将channel传入，并取出channel中的数据，以便下一步添加值的操作
	go recv(ch)

	ch <- true
	fmt.Println("main 函数结束")
	
}
