package main

import "fmt"

/*
	链式操作:
		- 通过结构体，实现 结构体.方法().方法()
*/

type student struct {
	name string
}

func (s student) learn() student{
	fmt.Printf("【%s】热爱学习\n", s.name)
	return s
}

func (s student) doHomework() student{
	fmt.Printf("【%s】热爱写作业\n", s.name)
	return s
}

func main() {

	tankStu := student{name: "tank jam"}

	// 普通调用
	//ret := tankStu.learn()
	//ret.doHomework()

	// 链式调用
	tankStu.learn().doHomework()

}
