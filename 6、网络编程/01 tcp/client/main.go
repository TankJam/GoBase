package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
	TCP CLIENT
*/



func main() {
	// 1.根据地址查找server并连接
	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("client connect to server failed, err:", err)
		return
	}
	defer conn.Close()

	for {
		// 2.向server端发送消息
		fmt.Println("请输入需要发送给服务端的消息: ")
		reader := bufio.NewReader(os.Stdin)  // 接收终端输入的内容
		input, err := reader.ReadString('\n')  // 从字符串中读取字节，直到遇到\n

		fmt.Printf("%T - %s", input, input)
		if err != nil{
			fmt.Println("recv right data failed,err: ", err)
			return
		}
		_, err = conn.Write([]byte(input))  // 将用户输入的内容，写入conn对象中，发送给服务端
		if err != nil{
			fmt.Println("send msg failed, err: ", err)
			return
		}

		if input == "q" || input == "Q"{
			// 输入q退出
			break
		}
	}
}
