package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	http server 提交form表单
*/

// search get请求渲染form表单页面
func search(w http.ResponseWriter, r *http.Request){
	data, err := ioutil.ReadFile("./6、网络编程/08 http server (form表单)/form.html")
	if err != nil{
		fmt.Println("read html file failed, err: ", err)
		return
	}
	w.Write(data)
}


// index 将用户输入的内容提交到服务端
func index(w http.ResponseWriter, r *http.Request){
	fmt.Println("请求的方法", r.Method)
	r.ParseForm()  // 解析form表单
	// 获取表单中的数据
	fmt.Printf("%#v\n", r.Form)
	usernameValue := r.Form.Get("username")
	passwordValue := r.Form.Get("pwd")
	fmt.Println(usernameValue, passwordValue)
	w.Write([]byte("index"))  // 提交玩返回index给页面
}

func main() {
	http.HandleFunc("/web", search)
	http.HandleFunc("/index", index)
	http.ListenAndServe("127.0.0.1:9527", nil)
}
