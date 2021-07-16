package main

import "fmt"

/*
	结构体内模拟 “继承”
*/

// 动物结构体
type animal struct{
	name string
}

// 定义一个动物 移动 的方法
func (a *animal) move(){
	fmt.Printf("%s moving~~~", a.name)
}

// 狗结构体
type dog struct{
	feet int
	animal  // 子结构体 继承 父结构体
}

// 定义一个狗 叫 的方法
func (d *dog) wang(){
	fmt.Printf("%s 在叫: 汪汪汪~~~", d.name)
}


func main() {
	var d1 = dog{
		feet:   4,
		animal: animal{
			name: "旺仔",
		},
	}

	d1.wang()  // dog 调用自己绑定的 wang() 方法
	d1.move()  // dog 调用了父结构体 animal 的 move() 方法
}
