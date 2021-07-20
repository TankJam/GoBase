package main

import "fmt"

/*
	switch 的坑
*/

func main() {
	f := func() bool{
		return false
	}

	// switch _ = f() ---> _ == false
	// 返回的是true
	switch _ = f(); {
	case false:
		fmt.Println("假")
	case true:
		fmt.Println("真")
	}

	a := false
	// 返回的是false
	switch a{
	case true:
		fmt.Println("真")
	case false:
		fmt.Println("假")

	}
}
