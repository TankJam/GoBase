package main

import (
	"bufio"
	"fmt"
	"net"
	"proto"
)

/*
	解决粘包现象 - server
		- 利用自己定义的协议 proto.Decode() 来拆解数据包，确定需要传输的数据量，解决粘包问题
*/

// process 服务端进程
func process(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	// 循环读取
	for {
		msg, err := proto.Decode(reader)  // 将数组，变成切片

		if err != nil{
			fmt.Println("decode msg failed, err: ", err)
			break
		}
		// 客户端发送了20次，但是每次服务端可以接收1024，导致数据黏在一起
		fmt.Println("收到 client 发送过来的数据: ", msg)
	}
}


func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9527")
	if err != nil{
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("accept failed, err: ", err)
			return
		}

		go process(conn)
	}
}
