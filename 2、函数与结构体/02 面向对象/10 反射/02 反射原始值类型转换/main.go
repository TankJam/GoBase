package main

import (
	"fmt"
	"reflect"
)

/*
	通过反射获取 "原始值，并强转为对应的数据类型";
*/

// reflectValue 通过反射获取类型信息
func reflectValue(x interface{}){
	// 获取x接口的值类型信息
	v := reflect.ValueOf(x)
	k := v.Kind()  // 获取值对应类型
	// 判断k
	switch k {
	case reflect.Int64:  // 值类型是否是Int64
		// v.Int() 获取反射得到的整型 原始值
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))

	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a)
	reflectValue(b)

	// 将int类型的原始值转为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c: %T\n", c)
}
