package main

import "fmt"

/*
	接口值: 由两部分组成，类型和值;
*/

func main() {

	var x interface{} // <type, value>
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1

	x = a  // <int64, 1000>
	x = b  // <int32, 10>
	x = c  // <int8, 1>
	x = 12.34  // <float64, 12.34>
	x = false  // <bool, false>
	fmt.Println(x)

	// 类型断言
	// 猜对了: 返回 true, 对应类型的值
	// 猜错了: 返回 false,对应的零值
	value, ok := x.(int)  // 判断x是否为int类型
	fmt.Printf("ok: %t  -  value:%#v  -  type:%T\n", ok, value, value)
}
