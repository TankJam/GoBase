package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	- 启动goroutine
	- 利用sync.WaitGroup 实现优雅等待
		- 计数器: 结束协程的标识符
*/

var wg sync.WaitGroup

// hello 定义一个hello函数
func hello(i int){
	defer wg.Done()  // 每执行完一个协程任务，则计数器 -1
	fmt.Println("hello tank jam!", i)

}

func main() {
	defer fmt.Println("哈哈哈 结束啦")

	// 设置GMP协程的 CPU 核数  m:n  10:2  10个任务2个CPU来处理
	runtime.GOMAXPROCS(1)  // 设置只使用1核来处理

	// 初始化计数器的值为10，可以允许10个协程执行
	wg.Add(10)  // 计数器 + 10

	// 创建10个协程
	for i := 0; i < 10; i++ {
		go hello(i)
	}

	fmt.Println("hello main init ...")

	// 等待hello执行完 (执行hello函数的goroutine执行完)
	wg.Wait()  // 阻塞态，等待所有goroutine执行结束
	fmt.Println("所有协程执行完毕，main函数结束...")

}
