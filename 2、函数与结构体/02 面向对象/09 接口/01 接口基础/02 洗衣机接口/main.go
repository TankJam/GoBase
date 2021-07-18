package main

import "fmt"

/*
	洗衣机接口
		- 实现 wash、dry功能
		- 让不同洗衣机厂商都统一使用同一个接口
*/

type xiYiJi interface {
	wash()
	dry()
}

// 海尔
type Haier struct {
	name, mode string
	price float64
}

// 田螺姑娘
type tianluo struct {
	name string
}

func (h Haier) wash(){
	fmt.Println("海尔洗衣机能洗衣服...")
}

func (h Haier) dry(){
	fmt.Println("海尔洗衣机自带甩干功能...")
}

func (t tianluo) wash(){
	fmt.Println("田螺姑娘可以洗衣服...")
}

func (t tianluo) dry(){
	fmt.Println("田螺姑娘可以把衣服拧干...")
}

func main() {
	var x xiYiJi  // 声明一个xiyiji类型的变量x
	h1 := Haier{
		name:  "小神童",
		mode:  "滚筒",
		price: 998.88,
	}
	fmt.Printf("%T\n", h1)

	// 所有赋值给接口的时候，必须绑定 xiYiJi 接口中存在的方法，相当于强制
	x = h1  // 接口是一种类型，一种抽象类型。
	fmt.Println(x)
	x.wash()
	x.dry()

	t1 := tianluo{name: "螺蛳粉"}
	x = t1
	fmt.Println(x)
	x.wash()
	x.dry()
}
