package main

import "fmt"

/*
	1.结构体指针操作
	2.何时使用指针类型:（构造函数 + 指针）    PS -- * 指针    & 接收者
		- 1）需要修改接收者中的值;
		- 2）接收者是拷贝代价比较大的对象;
		- 3）保证一致性，如果有某个方法使用了指针接收者，其他地方也应该使用指针接收者。
*/

// 定义结构体
type person struct{
	name, city string
	age int8
}

// 构造函数: 普通构造
func initPerson1(name, city string, age int8) person{
	// 返回结构体指针
	p := person{  // 实例化，产生了一次内存地址  x00
		name: name,
		city: city,
		age:  age,
	}

	fmt.Printf("value:  %#v  -- 普通构造实例化结构体地址: %p\n", p, p)

	return p
}

// 构造函数: 指针
func initPerson2(name, city string, age int8) *person{
	//  &person ---> *person
	p := &person{  // 对指针结构体的 接收者 进行实例化，指向的是同一个内存地址
		name: name,
		city: city,
		age:  age,
	}

	fmt.Printf("value:  %#v  -- 指针构造实例化结构体地址: %p\n", p, p)

	// 返回结构体指针
	return p
}


func main() {
	// 调用 new 产生一个空指针对象
	var p = new(person)  // --> 调用: *p --> 得到: &{}
	fmt.Println(p)

	// 1.指针对象字段名赋值
	//(*p2).name = "tank"
	//(*p2).city = "SH"
	//(*p2).age = 18
	p.name = "tank jam"
	p.city = "上海"
	p.age = 20
	fmt.Println(p)  // &{name city age}
	// 查看p2指针的类型
	fmt.Printf("%#v\n", p)

	// 2.取结构体的地址进行实例化
	p2 := &person{}  // --> 调用: &person{} --> 得到: *main.person
	p2.name = "kermit jam"
	p2.city = "广东广州"
	p2.age = 19
	fmt.Printf("%T\n", p2)  // 指针类型
	fmt.Printf("%#v\n", p2)

	fmt.Println("-------------------------------------------------------------------------")

	// 3.使用普通构造函数
	// 因为结构体是值传递，调用构造函数，得到的返回值是一个新的内存地址。
	p3 := initPerson1("tank zhang", "ZhanJiang", int8(20))
	fmt.Printf("value: %#v -- 调用构造函数的返回地址: %p\n", p3, p3)  // 调用几次就会拷贝几次内存地址

	fmt.Println("-------------------------------------------------------------------------")

	// 4.使用指针构造函数
	p4 := initPerson2("super tank", "广东省湛江市赤坎区", int8(22))
	fmt.Printf("value: %#v -- 调用指针构造函数的返回地址: %p\n", p4, p4)

}
