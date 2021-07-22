package main

import (
	"fmt"
	"math"
	"time"
)

/*
	select 多路复用
*/

// 创建两个string类型的通道，容量分别为100
var ch1 = make(chan string, 100)
var ch2 = make(chan string, 100)


// f1 每50毫秒往通道中添加数据
func f1(ch chan string){
	// math.MaxInt64 == 1<<63 - 1
	// == 9223372036854775807
	for i := 0; i < math.MaxInt64; i++{
		ch <- fmt.Sprintf("f1: %d", i)
		time.Sleep(time.Millisecond * 50)
	}
}

// f2 每100毫秒往通道中添加数据
func f2(ch chan string){
	for i := 0; i < math.MaxInt64; i++ {
		ch <- fmt.Sprintf("f2: %d", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	go f1(ch1)  // 往ch1这个通道中放f1开头的字符串
	go f2(ch2)  // 往ch2这个通道中放f2开头的字符串

	for {
		// 通过select实现去不同的通道取出值
		select {
		case ret := <- ch1:
			fmt.Println(ret)
		case ret := <- ch2:
			fmt.Println(ret)
		default:
			fmt.Println("暂时取不到数据!")
			time.Sleep(time.Millisecond * 500)
		}
	}

}
