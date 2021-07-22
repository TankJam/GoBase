package main

import (
	"fmt"
	"io"
	"net"
)

/*
	Http Client:
		- 模拟客户端，像服务端发送请求获取数据;
*/


func main() {
	conn, err := net.Dial("tcp", "www.liwenzhou.com:80")
	if err != nil{
		fmt.Println("访问失败!")
		return
	}

	defer conn.Close()

	// 发送数据: 发送请求头数据给目标节点
	conn.Write([]byte("GET /HTTP/1.0\r\n\r\n"))

	// 接收数据: 接收目标节点返回的数据
	var buf [1024]byte

	for {
		n, err := conn.Read(buf[:])

		if err == io.EOF{
			fmt.Print(string(buf[:n]))
			return
		}

		if err != nil{
			fmt.Println("接收数据失败, err: ", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
