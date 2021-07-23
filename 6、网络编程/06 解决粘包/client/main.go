package main

import (
	"fmt"
	"net"
	"proto"  // 自定义包需要放在当前项目根目录下的src目录中
)

/*
	解决粘包现象 - client
		- 利用自己定义的协议 proto.Encode() 来封装数据包，确定需要传输的数据量，解决粘包问题
*/

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	if err != nil {
		fmt.Println("dial failed, err: ", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++{
		msg := `Hello, Hello~ How are you?`

		// 封装数据包
		pkg, err := proto.Encode(msg)
		if err != nil{
			fmt.Println("encode msg failed, err: ", err)
			return
		}
		conn.Write(pkg)
	}
}
