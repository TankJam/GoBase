package main

import (
	"fmt"
	"net"
)

/*
	TCP SERVER:
		- net 网络包
*/

var msg chan string

// 单独处理连接的函数
func process(conn net.Conn, msg chan string){
	// 从连接中接收数据，每次只能接收1024个字节
	var buf [1024]byte
	n, err := conn.Read(buf[:])  // n表示读了多少数据
	if err != nil{
		fmt.Println("recv data from client failed, err: ", err)
	}
	// string(buf[:n])  将字节数据转为字符串
	fmt.Println(buf[:n])
	fmt.Printf("recv data from client, msg: [%s]", string(buf[:n]))

	if string(buf[:n]) == "q"{
		msg <- string(buf[:n])
	}
}


func main() {

	msg = make(chan string, 1)

	// 1.监听端口
	networkMode := "tcp"
	address := "127.0.0.1:9527"
	listener, err := net.Listen(networkMode, address)
	if err != nil{
		fmt.Println("listen:9527 failed, err: ", err)
		return
	}
	fmt.Printf("start server (%s)...", address)
	defer listener.Close()

	// 2. 接收客户端请求简历连接
	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("connect failed, err: ", err)
			continue
		}
		for {
			// 3.创建goroutine处理连接
			go process(conn, msg)
			ret := <- msg
			if ret == "q"{
				defer conn.Close() // 释放当前goroutine处理的连接
				break
			}
		}

	}
}
