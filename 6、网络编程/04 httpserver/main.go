package main

import (
	"fmt"
	"net/http"
)

/*
 	HTTP SERVER
*/

func sayHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "<h1 style='color:green'>tank jam</h1>")
}

func main() {
	// 注册路由
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9527", nil)  // 建立监听
	if err != nil{
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}
}
