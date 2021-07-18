package main

import "fmt"

/*
	什么是接口:
		- 接口就是一种类型，一种抽象的类型。
	为什么要用接口:
		- 为了将不同的类型存放在同一个切片中。

	接口是OOP中多态的表现形式，其实就是抽象类。
*/

// Cat 结构体
type Cat struct{}

// Say 猫叫
func (c Cat) Say() string{
	return "喵喵喵"
}

// Dog 结构体
type Dog struct{}

// Say 狗叫
func(d Dog) Say() string{
	return "汪汪汪"
}

// Pig 结构体
type Pig struct {}

// Say 猪叫
func (p Pig) Say() string{
	return "哼哼哼"
}


// 定义 动物 会叫的接口
// 目的是为了让不同的对象都能存放在同一个切片中
type animal interface {
	Say() string
}


func main() {
	// 普通实现: 让猪猫狗都叫
	p := Pig{}
	fmt.Println("pig: ", p.Say())
	c := Cat{}
	fmt.Println("cat: ", c.Say())
	d := Dog{}
	fmt.Println("dog: ", d.Say())


	// 接口实现
	// 猪猫狗都属于动物，所以都是用 animal 接口，接口中调用Say()方法
	var animalArray []animal
	p2 := Pig{}
	c2 := Cat{}
	d2 := Dog{}
	animalArray = append(animalArray, p2, c2, d2)
	fmt.Println(animalArray)
	for _, obj := range animalArray{
		ret := obj.Say()
		fmt.Println(ret)
	}
}
