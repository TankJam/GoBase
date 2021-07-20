package main

import (
	"fmt"
	"time"
)

/*
	内置的time包
*/

// timeStamp2timeObj 1.格式化时间函数
// interface{} 可以返回任意类型
func timeStamp2timeObj(timestamp int64)interface{}{  // timestamp接收时间戳
	timeObj := time.Unix(timestamp, 0)  // 将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	formatTime := fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d\n",
		year, month, day, hour, minute, second)
	return formatTime
}

// tickTime 定时器函数
func tickTime(){
	ticker := time.Tick(time.Second)  // 定义一个1秒间隔的定时器
	for i := range ticker{
		fmt.Println(i)
	}
}


func format2str(){
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15时4分
	//now.Format("格式化时间模板")
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}


func main() {
	//time.Time struct
	now := time.Now()
	// 获取当前时间对象
	fmt.Printf("$#v\n", now)
	fmt.Println(now.Year())  // 获取年
	fmt.Println(now.Month())  // 获取月
	fmt.Println(now.Day())  // 日
	fmt.Println(now.Hour())  // 时
	fmt.Println(now.Minute())  // 分
	fmt.Println(now.Second())  // 秒
	fmt.Println(now.Nanosecond())  // 获取当前时间的ns数

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())  // 获取时间戳纳秒数
	tm := now.Unix()

	// 格式化时间
	formatTime := timeStamp2timeObj(tm)
	fmt.Println(formatTime)

	// 定时器
	//tickTime()

	// 格式化时间 ---> 字符串格式
	format2str()
}
