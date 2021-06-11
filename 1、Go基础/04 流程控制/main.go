package main

import "fmt"

func main() {
	// 1. if、else if、else
	// 传统版
	// 判断 x 是否大于y的值
	y := 300
	x := 200
	if x > y {
		fmt.Println("x 大于 y")
	} else {
		fmt.Println("y 大于 x")
	}

	// 特殊版
	// 判断 z 是否大于y的值
	if z := 400; z > y {
		fmt.Println("z 大于 y")
	} else {
		fmt.Println("y 大于 z")
	}

	// 2.for 循环
	// 循环0-9的数字
	// 方式一:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------------------------------------------")
	// 方式二:
	i2 := 0
	for i2 < 10 {
		fmt.Println(i2)
		i2++
	}
	fmt.Println("------------------------------------------------------")

	// 无线循环 取代 while循环
	i3 := 0
	for {
		// 当i3大于10就结束循环
		if i3 > 10 {
			break // 退出循环
		} else {
			// 无线循环打印i3
			fmt.Println(i3)
		}

		i3++
	}
	fmt.Println("------------------------------------------------------")

	// 键值循环
	// for range(键值)  返回下标与值

	// 字符串
	str1 := "坦克非常强大"
	for key, value := range str1 {
		// %c  格式化中文
		fmt.Printf("key: %d, value: %x , %c \n", key, value, value)
	}

	for _, value := range str1 {
		// %c  格式化中文
		fmt.Printf("value: %x , %c \n", value, value)
	}

	fmt.Println("------------------------------------------------------")

	// 数组  值传递
	// 指定创建的长度
	ararrys := [4]int{1, 2, 3, 4}
	for _, value := range ararrys {
		fmt.Println(value)
	}

	fmt.Println("------------------------------------------------------")

	// 切片  引用传递
	// 无需指定长度，因为是指针，可以根据切片的值进行扩减
	var slice []int = []int{1, 2, 3, 4}
	for _, value := range slice {
		fmt.Println(value)
	}
	fmt.Println("------------------------------------------------------")

	// map map[key类型]值类型
	var m = map[string]int{
		"height": 300,
		"width":  400,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("------------------------------------------------------")

	// 遍历 channel 管道  make(类型)  自动扩容
	c := make(chan int) // 初始化一个能插入整型的channel管道
	// 启动协程
	go func() {
		c <- 1 // 往c管道插入1
		c <- 2 // 往c管道插入2
		c <- 3 // 往c管道插入3
		c <- 4 // 往c管道插入4
		c <- 5 // 往c管道插入5
		close(c)
	}() // 调用匿名函数

	// 从管道中取出一个个值
	//num1 := <- c
	//num2 := <- c
	//num3 := <- c
	//num4 := <- c
	//num5 := <- c
	//fmt.Println(num1, num2, num3, num4, num5)

	// 遍历管道中的数据
	for v := range c {
		fmt.Println(v)
	}

}
