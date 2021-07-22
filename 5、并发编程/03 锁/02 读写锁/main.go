package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	读写锁: 读的操作比写的操作多的时候使用读写锁，能提高性能。
		- 测试普通测试锁的执行时间
		- 测试读写锁的执行时间
			- 读锁:
				- wrLock.RLock()
				- wrLock.RUnlock()
			- 写锁:
				- wrLock.Lock()
				- wrLock.Unlock()
*/

var x int64
var wg sync.WaitGroup
var lock sync.Mutex  // 互斥锁
// 读写锁:多个goroutine同时读加的是读锁，写的时候是写锁
var rwLock sync.RWMutex

// read 读操作
func read(){
	defer wg.Done()
	//// 测试使用互斥锁
	//lock.Lock()
	//fmt.Println("read ...")
	//time.Sleep(time.Millisecond * 1) // 模拟每次读操作耗费1毫秒
	//lock.Unlock()

	// 测试使用读写锁
	rwLock.RLock()
	time.Sleep(time.Millisecond * 1)
	rwLock.RUnlock()
}

// write 写操作
func write(){
	defer wg.Done()
	//// 测试使用互斥锁
	//lock.Lock()
	//fmt.Println("write ................")
	//time.Sleep(time.Millisecond * 5) // 模拟每次写操作耗费5毫秒
	//lock.Unlock()

	// 测试使用读写锁
	rwLock.Lock()
	time.Sleep(time.Millisecond * 1)
	rwLock.Unlock()
}

func main() {
	// 计算读写时间
	start := time.Now()

	// 写10次任务
	for i := 0; i < 10; i++{
		wg.Add(1)
		go write()
	}

	// 读1000次任务
	for i := 0; i < 1000; i++{
		wg.Add(1)
		go read()
	}

	wg.Wait()

	end := time.Now()

	//end.Sub() 时间戳与时间戳相减可以通过该方法实现
	fmt.Printf("耗费了: [%v] 时间~", end.Sub(start))
}
