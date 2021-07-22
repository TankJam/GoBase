package fib

import "testing"

/*
	fib 测试用例
*/

// Fib 基准测试
func BenchmarkFib(b *testing.B) {
	// 连接数据库
	// b.ResetTimer()  // 清除时间
	for i := 0; i < b.N; i++{
		Fib(2)
	}
}

func main() {
	
}
