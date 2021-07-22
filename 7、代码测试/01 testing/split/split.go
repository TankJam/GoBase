package split

import (
	"fmt"
	"strings"
)

/*
	定义一个切割字符串的包
*/

// Split 用sep分割s
// a:b:c : --> [a b c]
func Split(s, sep string)[] string{
	count := strings.Count(s, sep)  // 数一下字符串s中包含多少个sep
	result := make([]string, 0, count + 1)  // 根据sep的数量初始化切片
	index := strings.Index(s, sep)  // 查看第一次出现的sep的index位置
	//fmt.Println(index)
	for index >= 0 {
		// 将 sep 左边的数据都追加到切片中
		result = append(result, s[:index])
		// 把 sep 右边的数据重新赋值给s
		s = s[index + len(sep): ]
		// 再次获取 rep 的index, 直到 s 字符串中的 sep不再出现，自动退出循环
		index = strings.Index(s, sep)
	}
	// 将最后一个 sep 右边的值追加到切片中
	result = append(result, s)
	return result
}


// Add 实现字符串相加的函数
func Add(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}