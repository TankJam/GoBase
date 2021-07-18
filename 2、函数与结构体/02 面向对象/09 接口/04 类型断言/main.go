package main

import "fmt"

/*
	类型断言
		- type nullInterface interface {} // 太繁琐
*/

type Cat struct {}


//showType 可接收任意类型变量的函数
func showType(x interface{}){
	// 类型断言
	v1, ok := x.(int)  // 猜测x是否为int类型
	if !ok{
		// 猜错了
		fmt.Println("x 不是 int")
	}else{
		fmt.Println("x 是 int类型", v1)
	}

	v2, ok := x.(string) // 猜测x是否为string类型
	if !ok{
		fmt.Println("x 不是 string")
	}else{
		fmt.Println("x 是 string类型", v2)
	}
}

// justifyType switch版本 类型断言
func justifyType(x interface{}){
	switch v := x.(type){  // 断言判断x的类型
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int, value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	case Cat:
		fmt.Printf("x is a Cat, value is %v\n", v)
	case *string:
		fmt.Printf("x is a string poninter, value is %v\n", v)
	default:
		fmt.Println("un support type!")
	}
}

func main() {
	var x interface{}
	x = 100
	showType(x)

	// switch版本 类型断言
	justifyType(100)
	justifyType("tank jam")
	s := "string poninter"
	justifyType(&s)
	justifyType(Cat{})
}
