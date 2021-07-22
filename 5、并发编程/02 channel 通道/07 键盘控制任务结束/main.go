package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 生产者消费者模型
// 使用goroutine和channel实现一个简易的生产者消费者模型

// 生产者：产生随机数  math/rand

// 消费者：计算每个随机数的每个位的数字的和     14134134123 = ?

// 1个生产者 20个消费者

var itemChan chan *item
var resultChan chan *result
var exitChan chan struct{}  // 用于检测是否有用户按键盘退出的通道

type item struct {
	id  int64
	num int64
}

type result struct {
	item *item
	sum  int64
}

// 生产者
// chan *item :既能接收也能发送的一个队列
// chan<- *item :只能发送的一个队列（单向通道）
// <-chan *item :只能接收的一个队列（单向通道）
func producer(ch chan<- *item) {
	// 1. 生成随机数
	var id int64
	for {
		id++
		number := rand.Int63() // int64正整数
		tmp := &item{
			id:  id,
			num: number,
		}
		// 2. 把随机数发送到通道中
		ch <- tmp
	}
}

// 计算一个数字每个位的和
func calc(num int64) int64 {
	// 123%10=12...3  sum = 0 + 3
	// 12%10=1...2
	// 1%10=0...1
	var sum int64 // 0
	for num > 0 {
		sum = sum + num%10 // sum = 5 + 1
		num = num / 10     // num = 0
	}
	return sum
}

// 消费者
func consumer(ch chan *item, resultChan chan *result) {
	for tmp := range ch {
		// (*tmp).num // item.num
		sum := calc(tmp.num)
		// 构造result结构体
		retObj := &result{
			item: tmp,
			sum:  sum,
		}
		resultChan <- retObj
	} // 结构体指针 *item
}

func startWorker(n int, ch chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultChan)
	}
}

// 打印结果
func printResult(exitChan chan struct{}, resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("id:%v, num:%v, sum:%v\n", ret.item.id, ret.item.num, ret.sum)
		//time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	itemChan = make(chan *item, 10000)
	resultChan = make(chan *result, 10000)
	exitChan = make(chan struct{}, 1)

	go producer(itemChan)  // 生产随机数

	startWorker(20, itemChan, resultChan)  // 消费随机数

	go func() {  // 获取键盘输入
		// os.Stdin: *os.File  获取输入的文件数据
		os.Stdin.Read(make([]byte, 1))
		exitChan <- struct{}{}  // 用户敲了键盘表示要退出
	}()

	// 打印结果
	go printResult(exitChan, resultChan)

	// select会等待，等到exitChan通道中有值时退出
	select {  // 循环多路监听
	case <- exitChan:
		break
	}
}