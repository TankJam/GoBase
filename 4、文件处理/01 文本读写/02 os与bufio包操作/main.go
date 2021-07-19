package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	bufio读取数据: 利用缓冲区从文件读取数据
*/

//readByLine bufio读数据
func readByLine(){
	file, err := os.Open("./4、文件处理/01 文本读写/02 os与bufio包操作/xx.txt")
	if err != nil{
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()

	// 利用缓冲区从文件读取数据
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')  // 字符
		if err == io.EOF{
			fmt.Print(str)
			return
		}
		if err != nil{
			fmt.Println("读取文件内容失败!")
			return
		}
		fmt.Print(str)
	}
}

func main() {
	readByLine()
}

