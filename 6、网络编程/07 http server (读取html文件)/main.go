package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	http server 加载html
*/

func sayHello(w http.ResponseWriter, t *http.Request){
	// 从hello.html文件中读取数据写入到w中，并渲染到浏览器中
	data, err := ioutil.ReadFile("./6、网络编程/07 http server (读取html文件)/hello.html")
	if err != nil {
		fmt.Println("read from file failed, err: ", err)
		return
	}
	w.Write(data)  // 渲染页面
}

func main() {
	http.HandleFunc("/", sayHello)
	// run server
	err := http.ListenAndServe("127.0.0.1:9527", nil)
	if err != nil {
		panic(err)
	}
}
