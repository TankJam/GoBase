package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	原子操作: atomic.Add()
		- 1.并发下保证数据安全
		- 2.性能优于枷锁

	- 内置整数支持的操作
*/


var x int64
var l sync.Mutex
var wg sync.WaitGroup

// 给x做赋值操作
// 不加锁: 线程不安全
func add(){
	x++
	wg.Done()
}

// 加锁: 线程安全，但性能差
func mutexAdd(){
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}

// 原子操作: 线程安全，并且性能比加锁要高
func atomicAdd(){
	atomic.AddInt64(&x, 1)  // x++
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000000; i++{
		wg.Add(1)
		//go add()  // 普通版
		//go mutexAdd()  // 加锁版
		go atomicAdd()  // 原子版
	}
	wg.Wait()
	fmt.Println(x)
	end := time.Now()
	fmt.Printf("消耗时间为: %s", end.Sub(start))
	
}
