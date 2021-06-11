package main

import (
	"fmt"
	"math"
)

func main() {
	// 1.整型
	age := 18
	println(age)

	// 2.浮点型
	pi := math.Pi // 获取 圆周率
	fmt.Printf("%f\n", pi)
	fmt.Printf("%.2f\n", pi) // 只要六两位小数

	// 3.复数类型 (暂时不理解他的应用场景)
	var c1 complex64
	var c2 complex128
	c1 = 1 + 2i
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// 4.字符串类型

	// 5.布尔类型
	// 6.byte 与 rune
	// 遍历字符串
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	// 7.类型转换
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}
