package main

import (
	"fmt"
	"os"
)

/*
打开文件支持文件写入
	- file.Write([]byte("字符"))  写入字符前需要先转成字节
	- file.WriteString(str)   写入字符串
*/

func main() {
	// O_APPEND 追加   O_WRONLY 支持写   O_CREATE 创建
	file, err := os.OpenFile("./5、文件处理/01 文本读写/04 写文件/xx.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil{
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()

	str := "tank jam is very handsome!!!\n"
	file.Write([]byte("嘿嘿嘿\n"))  // 写入嘿嘿嘿与换行的字节
	file.WriteString(str)  // 写入字符串
}
