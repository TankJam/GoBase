package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// trace的编程过程
// 1.创建trace文件
// 2.启动
// 3.停止

func main() {
	// 1.创建trace文件
	f, err := os.Create("trace.out")
	if err != nil{
		panic(err)
	}

	defer f.Close()

	// 2.启动
	err = trace.Start(f)
	if err != nil{
		panic(err)
	}
	fmt.Println("hello GMP")

	// 3.停止
	trace.Stop()

}
