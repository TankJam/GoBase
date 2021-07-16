package main

import "fmt"

/*


Go语言中 函数 与 方法 的区别:
	- 函数: 谁都可以调用;
	- 方法: 绑定给某个具体的类型使用的;
		- 定义: 函数名前指定接受者之后就是方法
		- 注意: Go语言中没有self与this，而是规定成俗使用类型的首字母小写
*/

// 结构体定义
type people struct{
	name, gender string
}


// people的方法
//func dream(){}  // 函数
//func (类型首字母小写 *类型) dream(){}  // 方法
func (p *people) dream(){
	p.gender = "male"
	fmt.Printf("%s的梦想是不用上班也有收入，可以环游世界~\n", p.name)
}


// 可以给任意类型追加方法, 但不能给别的包定义的类型添加方法
type MyInt int

func (m *MyInt) sayHi(){
	fmt.Println("hello MyInt...")
}

func main() {
	var p = &people{
		name:   "tank jam",
		gender: "男",
	}
	fmt.Println(p.gender)
	p.dream()
	fmt.Println(p.gender)

	// 给任意类型追加方法
	var a MyInt
	fmt.Println(a)
	a.sayHi()
}

