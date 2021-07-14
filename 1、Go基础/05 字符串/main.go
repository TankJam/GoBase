package main

import (
	"fmt"
	"strings"
)

/*
	- 字符串常用操作
		- 获取长度：
			- len
		- 字符串拼接:
			- 字符串 + 字符串
			- fmt.Sprintf
   		- 字符串切割
			- strings.Split()
		- 是否包含
			- strings.Contains()
		- 判断开头
			- strings.HasPrefix
		- 判断结尾
			- strings.HasSuffix
		- 判断字符串位置 开头
			- strings.Index()
		- 判断字符串位置 最后
			- strings.LastIndex()
		- 列表的字符串拼接
			- strings.Join()
*/
func main() {
	s := "hello tank"
	// 1.求字符串的长度
	fmt.Println(len(s))

	// 2.拼接字符串
	s2 := "hello 小豆丁"
	fmt.Println(s + " - " + s2)
	s3 := fmt.Sprintf("%s - %s\n", s, s2)  // 格式化拼接
	fmt.Println(s3)

	// 3.字符串切割,切割后会将切割后的字符串追加到切片中
	s4 := "how do you do"
	slice := strings.Split(s4, " ")
	fmt.Println(slice)

	// 4.判断是否包含 Contains
	isStr := "how"
	flag := strings.Contains(s4, isStr)
	fmt.Println(flag)
	if flag{
		fmt.Printf("包含: %s\n", isStr)
	}else{
		fmt.Printf("不包含: %s\n", isStr)
	}

	// 5.判断前缀 strings.HasPrefix()
	fmt.Println(strings.HasPrefix(s4, "how"))  // true

	// 6.判断后缀 strings.HasSuffix()
	fmt.Println(strings.HasSuffix(s4, "do"))  // true

	// 7.判断字符串的位置（获取第一个字符串出现的位置）
	fmt.Println(strings.Index(s4, "o"))

	// 8.字符串最后出现的位置
	fmt.Println(strings.LastIndex(s4, "o"))

	// 9.Join 数组/切片中字符串的拼接
	slice2 := []string{"how", "do", "you", "do"}
	fmt.Println(slice2)
	s5 := strings.Join(slice2, " ")
	fmt.Println(s5)
}
