package main

import (
	"fmt"
	"time"
)

/*
time包练习题
	1. 编写一个函数，接收Time类型的参数，函数内部将传进来的时间格式化输出为`2017/06/19 20:30:05`格式。
	2. 编写程序统计一段代码的执行耗时时间，单位精确到微秒。
*/

func formatTime(t time.Time){
	timeStr := t.Format("2006/06/02 15:04:05")
	fmt.Println(timeStr)
}

// 计算执行时间，以微妙来计算
func calcTime(){
	start := time.Now()
	fmt.Println(start.UnixNano())  // 纳秒
	startTimeStamp := start.UnixNano() / 1000  // 得到微秒
	fmt.Println("围城啊，围个了城呀呀呀！")
	time.Sleep(time.Millisecond * 30)  // 等待30毫秒  1000毫秒 == 1秒
	end := time.Now()
	endTimeStamp := end.UnixNano() / 1000
	fmt.Printf("消耗了%d微秒\n", endTimeStamp - startTimeStamp)
}

func calcTime2(){
	start := time.Now()
	fmt.Println("围城啊，围个了城呀呀呀！")
	time.Sleep(time.Millisecond * 30)
	fmt.Println("消耗了: ", time.Since(start))
}



func main() {

	calcTime()

	calcTime2()

}
