package main

import (
	"fmt"
	"io/ioutil"
)

/*
	ioutil读取数据:
*/


//readFile ioutil读取数据
func readFile(filename string){
	// 内部会调用 defer f.Close() 与 相当于python的with
	content, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("read file failed, err: ", err)
		return
	}
	fmt.Println(content)
	fmt.Println(string(content))
}


func main() {
	readFile("./4、文件处理/01 文本读写/03 os与ioutil操作/xx.txt")
}
