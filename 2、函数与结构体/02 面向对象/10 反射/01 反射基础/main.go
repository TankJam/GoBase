package main

import (
	"fmt"
	"reflect"
)

/*
	反射TypeOf(): 通过反射获取类型信息
*/

func reflectType(x interface{}){
	t := reflect.TypeOf(x) // 动态获取x的类型信息
	fmt.Printf("type: %v\n", t)

	fmt.Printf("name: %v  -  kind:%v\n", t.Name(), t.Kind())

	// 注意: 指针的 t.Name() 是空的
}

type cat struct{
	name string
}

type person struct {
	name string
	age uint8
}

func main() {
	reflectType(100)
	reflectType(false)

	x := 100
	reflectType(&x)  // // 注意: 指针的 t.Name() 是空的

	// 测试自定义的结构体类型
	var c1 = cat{
		name: "花花", // 有的人看起来人模狗样儿的，可是Ta连个猫都没有。
	}
	var p1 = person{
		name: "豪杰",
		age:  18,
	}
	reflectType(c1)
	reflectType(p1)

	var i int32 = 100
	var f float32 = 12.34
	reflectType(&i)
	reflectType(&f)

	var a = []int{1, 2, 3}
	reflectType(a)
	var b = [3]int{1, 2, 3}
	reflectType(b)
	reflectType(map[string]int{})
}
