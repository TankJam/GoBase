package main

import "fmt"

func main() {
	// 1.数组定义
	//var a[3]int
	//var b[3]string

	// 2.数组初始化
	// 方式一: 固定长度
	var a [3]int = [3]int{1, 2, 3}
	//var a[3]int = [3]int{1, 2, 3, 4}  // 报错，数组初始化是固定长度的，超过长度则报错
	fmt.Println(a)

	// 方式二: 编译器根据初始化的长度来自动推断数组的长度
	var b = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 方式三: 根据指定的索引进行初始化
	var c = [...]int{1: 1, 3: 3, 5: 5, 7: 7, 9: 9}
	fmt.Println(c)

	// 3.遍历数组
	var ararrys = [...]string{"tank", "广东", "大帅哥", "喜欢哈哈哈哈哈"}
	for _, v := range ararrys {
		fmt.Println(v)
	}

	// 4.二维数组
	// 注意： 多维数组只有第一层可以使用...来让编译器推导数组长度
	e := [...][2]string{
		{"上海", "长宁区"},
		{"北京", "海淀区"},
		{"广州", "天河区"},
	}

	fmt.Println(e)

	// 二维数组遍历
	// 先循环外层
	for _, v1 := range e {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		fmt.Println("========")
	}
}
