package main

import "fmt"

/*
	实现接口时，使用指针接收者和使用值接收者的区别
*/

type animal interface {
	speak()
	move()
}

type cat struct{
	name string
}

// cat 类型实现animal的接口
// 1.使用值接收者
func (c cat)speak(){
	fmt.Println("喵喵喵")
}

func (c cat)move(){
	fmt.Println("猫在动")
}

// 2.使用指针接收者
func (c *cat)speak2(){
	fmt.Println("喵喵喵")
}

func (c *cat)move2(){
	fmt.Println("猫在动")
}

func main() {
	var x animal

	h1 := cat{"花花"}
	x = h1  // 实现animal接口的是*cat类型，hh此时是一个cat类型
	x.speak()  // h1.speak()
	x.move()  // h1.move()
	fmt.Println(x)
	fmt.Printf("%T --- %p\n", x, x)

	tom := &cat{"汤姆"}
	x = tom
	x.speak()  // (*tom).speak()
	x.move()  // (*tom).move()
	fmt.Println(x)
	fmt.Printf("%T --- %p\n", x, x)
}
