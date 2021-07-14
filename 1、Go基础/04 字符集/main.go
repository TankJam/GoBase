package main

import "fmt"

/*
	字符集
*/
func main() {
	// byte: uint8的别名  ASCII码
	// rune: int32的别名  Unicode码
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	// %T 是打印类型
	fmt.Printf("c1:%T  c2:%T\n", c1, c2)

	s := "hello 坦克"
	for i := 0; i < len(s); i++ {
		// 打印s字符串的字符集
		fmt.Printf("%c\n", s[i])
	}
}
