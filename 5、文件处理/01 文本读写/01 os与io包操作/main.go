package main

import (
	"fmt"
	"io"
	"os"
)

/*
	打开和关闭文件
		- os.Open("文件目录") 打开文件

	注意: 路径 “./代表的项目工程目录，而不是当前文件的上一级目录”
*/

func main() {
	//// 路径: './xx.txt'
	////file, err := os.Open("./xx.txt")
	//file, err := os.Open("./5、文件处理/01 文本读写/01 打开与关闭文件/xx.txt")
	//if err != nil{
	//	fmt.Println("open file failed, err: ", err)
	//	return
	//}
	//
	//defer file.Close()  // 文件处理后关闭文件
	//
	//// 读文件
	//var tmp [128]byte  // 定义一个字节数组
	//// var s = make([]byte, 0, 128)
	//for {
	//	// file.Read(tmp[:])  将读取出来的数据，存放到tmp切片中
	//	n, err := file.Read(tmp[:])  // 基于数组得到一个切片
	//	if err == io.EOF{
	//		fmt.Println("文件已读完!")
	//		return
	//	}
	//	if err != nil{
	//		fmt.Println("read from file failed, err: ", err)
	//		return
	//	}
	//	fmt.Printf("读取了%d个字节\n", n)
	//	fmt.Println(string(tmp[:]))
	//}

	fileObj, err := os.Open("./5、文件处理/01 文本读写/01 os与io包操作/xx.txt")
	if err != nil{
		fmt.Println("open file failed, err: ", err)
		return
	}

	defer fileObj.Close()  // 关闭文件

	var tmp [128]byte  // 定义一个容量为128的字节数组

	for {
		n, err := fileObj.Read(tmp[:])  // 切片可以修改
		if err == io.EOF{
			fmt.Println("文件已经读取完毕!")
			return
		}
		if err != nil{
			fmt.Println("read from file failed, err: ", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(tmp[:]))  // 将字节类型，转为string
	}
}
