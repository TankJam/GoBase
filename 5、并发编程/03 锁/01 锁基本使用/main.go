package main

import (
	"fmt"
	"sync"
)

/*
	互斥锁: 加锁牺牲性能保证数据安全。
*/

var x int64 // 定义全局变量x
var wg sync.WaitGroup

// 定义一个互斥锁
var lock sync.Mutex

// add 定义一个函数，对全局变量x做累加的操作
func add(){
	for i := 0; i < 5000; i++ {
		// 抢锁，执行+1，把结果赋值给x写入内存中
		lock.Lock()  // 上锁
		x = x + 1
		lock.Unlock()  //释放锁，让下一个任务过来抢锁
	}

	wg.Done()

}


func main() {
	wg.Add(2)  // 设置数值为2的计数器
	// 启动两个任务，让他们抢锁并执行
	go add()
	go add()
	wg.Wait()  // 等待计数器为0，协程任务全部完成，则结束
	fmt.Println(x)
}
