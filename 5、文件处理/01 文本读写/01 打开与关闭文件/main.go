package main

import (
	"fmt"
	"io"
	"os"
)

/*
	打开和关闭文件
		- os.Open("文件目录") 打开文件
*/

func main() {

	file, err := os.Open("/Users/tank/Desktop/Go复习/GoBase/5、文件处理/01 文本读写/01 打开与关闭文件/xx.txt")
	if err != nil{
		fmt.Println("open file failed, err: ", err)
		return
	}

	defer file.Close()  // 文件处理后关闭文件

	// 读文件
	var tmp [128]byte  // 定义一个字节数组
	// var s = make([]byte, 0, 128)
	for {
		n, err := file.Read(tmp[:])  // 基于数组得到一个切片
		if err == io.EOF{
			fmt.Println("文件已读完!")
			return
		}
		if err != nil{
			fmt.Println("read from file failed, err: ", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(tmp[:]))
	}
}
