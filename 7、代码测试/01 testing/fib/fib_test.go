package fib

import "testing"

/*
	fib 测试用例
		- 压力测试
*/


/*
goos: windows
goarch: amd64
cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
BenchmarkFib
BenchmarkFib-8   	235710169	         4.731 ns/op
PASS
*/
// Fib 基准测试
func BenchmarkFib(b *testing.B) {
	// 连接数据库
	// b.ResetTimer()  // 清除时间
	for i := 0; i < b.N; i++{
		Fib(2)
	}
}


// 内部调用的函数
func benchmarkFib(b *testing.B, n int) {
	for i :=0; i < b.N; i++{
		Fib(n)
	}
}

/*
BenchmarkFib2
BenchmarkFib2-8   	254732937	         4.696 ns/op
*/
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

/*
BenchmarkFib3
BenchmarkFib3-8   	   31328	     38405 ns/op
*/
func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 20)
}
