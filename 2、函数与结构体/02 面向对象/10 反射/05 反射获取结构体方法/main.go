package main

import (
	"fmt"
	"reflect"
)

/*
	通过反射得到结构体的方法
*/

type Student struct{
	name string
	score int
}

// 给Student结构体添加 Study 与 Sleep 方法
func (s Student) Study(day int)string{
	msg := "好好学习，天天向上!"
	fmt.Println(msg)
	return msg
}

func (s Student) Sleep()string{
	msg := "好好睡觉, 快乐成长!"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}){
	t := reflect.TypeOf(x)  // 获取对象名称空间
	v := reflect.ValueOf(x)  // 获取对象的值信息

	fmt.Println(t.NumMethod())  // 获取方法数
	//fmt.Println(v.NumMethod())  // 获取方法数
	for i := 0; i < v.NumMethod(); i++{
		// v.Method(i) 根据i获取结构体方法
		methodType := v.Method(i).Type()  // 获取方法的类型
		fmt.Printf("method name: %s\n", t.Method(i).Name)
		fmt.Printf("method: %s\n", methodType)

		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		//var args = []reflect.Value{}
		//v.Method(i).Call(args)
	}
}

func main() {
	var stu1 = Student{
		name:  "Tank jam",
		score: 98,
	}
	printMethod(stu1)
}
