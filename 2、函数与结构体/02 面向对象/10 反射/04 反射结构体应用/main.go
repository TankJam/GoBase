package main

import (
	"fmt"
	"reflect"
)

/*
	结构体与反射的应用;
		- t := reflect.TypeOf()  返回传入变量的类型名称空间
			- 类似于Python的 obj.__dict__
		- t.NumField() 获取字段的数量
		- field := t.Field(i) 获取弟i个字段
			- field.Tag.Get("json")  // 获取可见性标签名
		- t.FieldByName("Score")  通过字段名获取字段信息
			- 类似于python的 getattr("Score")
			- 返回 值, true|false
*/

// Student
type Student struct {
	Name string `json:"name" hj:"mz"`
	Score int `json:"score" hj:"fs"`
}

func main() {
	stu1 := Student{
		Name:  "tank jam",
		Score: 99,
	}

	t := reflect.TypeOf(stu1)  // 返回类型对象

	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Tag.Get("hj"))  // 获取字段的可见hj标签名
		fmt.Println(field.Tag.Get("json"))  // 获取字段的可见hj标签名
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok{
		fmt.Printf("name: %s - index: %d - type:%v - json tag: %v - value:%v\n",
			scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"), scoreField)
	}
}
