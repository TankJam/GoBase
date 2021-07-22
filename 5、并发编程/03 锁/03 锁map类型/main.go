package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
	sync.Map:
		由于Go内置的map类型不是线程安全的，
		所以sync包提供了 sync.Map 是线程安全的。

	并发场景下推荐使用 sync.Map

执行: 会报错，因为map不是线程安全的。
*/

//// --------------- Go 内置 map 类型 ---------------
//var m = make(map[string]int)
//
//// get 根据key取值
//func get(key string) int {
//	return m[key]
//}
//
//// set 添加键值对或修改key对应的值
//func set(key string, value int){
//	m[key] = value
//}
//
//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i< 20; i++{
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)  // 将整型转为字符串
//			set(key, n)  // 给map设置键值对
//			fmt.Printf("k=:%v, v:=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}
//
//// -------------------------------------------


// --------------- sync.Map{} ---------------
var m = sync.Map{}
/*
	m.Store  添加键值对
	m.Load   根据key取值
*/

func tickDemo(){
	ticker := time.Tick(time.Second)  // 定义一个1秒间隔的定时器
	for i := range ticker{
		fmt.Println(i)  // 每秒都会执行的任务
	}
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i< 20; i++{
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)  // 将整型转为字符串
			m.Store(key, n)  // 往map中添加键值对
			value, _ := m.Load(key)  // 从sync.Map中取值
			fmt.Printf("k=:%v, v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
// -------------------------------------------