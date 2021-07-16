package main

import "fmt"

/*
	匿名字段
*/

type student struct{
	name string
	string  // 匿名字段1
	int  // 匿名字段2
}

func main() {
	stu1 := student{
		name:   "tank jam",
	}
	fmt.Println(stu1.name)
	fmt.Println(stu1.string)  // 匿名字段会打印默认值
	fmt.Println(stu1.int)
}
