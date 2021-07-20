package main

import (
	"fmt"
	"reflect"
)

/*
	通过反射修改值
*/

func modifyValue(x interface{}){
	v := reflect.ValueOf(x)  // reflect.Value
	kind := v.Kind()  // 获取值对应的类型
	fmt.Println(kind)
	if kind == reflect.Ptr{ // Ptr == pointer 判断是kind是否为指针
		// 若是指针，则修改值
		// 先通过reflect.ValueOf(x).Elem()获取元素值，然后再通过SetInt()修改值
		v.Elem().SetInt(2000)
	}
}

func main() {
	var a int64 = 200
	modifyValue(a)
	fmt.Println(a)

	modifyValue(&a)
	fmt.Println(a)
}
