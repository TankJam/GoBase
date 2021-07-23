package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

/*
	粘包现象 - server
*/

// process 服务端进程
func process(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte // 定义长度为1024的byte数组
	// 循环读取
	for {
		n, err := reader.Read(buf[:])  // 将数组，变成切片
		if err == io.EOF {
			break
		}
		if err != nil{
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		// 客户端发送了20次，但是每次服务端可以接收1024，导致数据黏在一起
		fmt.Println("收到 client 发送过来的数据: ", recvStr)
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
