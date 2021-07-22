package fib

/*
	Fib 是一个计算弟n个斐波那契数列的函数
	- 定义测试函数
*/

// Fib 是一个计算弟n个斐波那契数列的函数
func Fib(n int) int {
	if n < 2{
		return n
	}
	return Fib(n - 1) + Fib(n - 2) // 递归
}


