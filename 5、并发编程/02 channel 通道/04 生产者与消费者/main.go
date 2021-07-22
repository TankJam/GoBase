package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
生产者与消费者模型:
	- 使用goroutine和channel实现一个简易的生产者与消费者模型
		- 生产者: 随机生产数
			- math/rand包实现
		- 消费者: 计算每个随机数的每个位的数字的和

		- 1个生产者
		- 20个消费者

	- 创建检测 生产者 与 消费者 关闭的goroutine
*/

type item struct {
	id, num int64
}

type result struct{
	item *item
	sum int64
}

var itemChan chan *item  // 存放生产者数据
var resultChan chan *result  // 存放消费者数据
var wg sync.WaitGroup

var doneChan chan struct{}  // 接收结构体类型的chan，用于检测消费者执行完毕时关闭通道

// producer 生产者
func producer(ch chan *item){
	// 1.生成随机数
	var id int64
	for i := 0; i < 10000; i++{
		id++
		number := rand.Int63()  // 随机生成int64正整数
		tmp := &item{
			id:  id,
			num: number,
		}

		// 2.将随机数发送到通道中
		ch <- tmp
	}

	close(ch)
}

// calc 计算一个数字每个位的和
func calc(num int64) int64{
	var sum int64
	for num > 0{
		sum = sum + num % 10  // num = 5 + 1
		num = num / 10
	}
	return sum
}

// consumer 消费者
func consumer(ch chan *item, resultChan chan *result){
	defer wg.Done()
	for tmp := range ch{
		sum := calc(tmp.num)

		// 构造result结构体
		retObj := &result{
			item: tmp,
			sum:  sum,
		}

		resultChan <- retObj
	}  // 结构体指针 *tiem
}

// startWorker 开始消费
func startWorker(n int, ch chan *item, resultChan chan *result){
	for i := 0; i < n; i++{

		// 消费者对生产者的队列进行消费
		go consumer(ch, resultChan)
	}
}

// printResult 打印结果
func printResult(resultChan chan *result){
	for ret := range resultChan{
		fmt.Printf("id:%v, num:%v, sum:%v\n",
			ret.item.id, ret.item.num, ret.sum)
	}
	// 不知道什么时候应该关闭resultChan
	// 当20个消费者的goroutine都执行完毕的时候 要关闭resultChan
}

func main() {
	itemChan = make(chan *item, 10000)
	resultChan = make(chan *result, 10000)
	go producer(itemChan)
	wg.Add(20)
	startWorker(20, itemChan, resultChan)

	wg.Wait()  // 等待所有生产result的goroutine都结束再打印

	close(resultChan)

	printResult(resultChan)
}
