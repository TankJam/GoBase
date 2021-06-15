package main

import (
	"fmt"
	"sort"
)

// 数组的局限性，因为是限长，无法传递不同长度的参数给函数使用

// 1.只能接收长度为3的数组
func arraySum(x [3]int)int{
	sum := 0
	for _, v := range x{
		sum += v
	}
	return sum
}


func main() {
	// 数组的劣势
	var arrays = [3]int{1, 2, 3}
	var res = arraySum(arrays)
	fmt.Println(res)

	//arrays2 := [4]int{1, 2, 3, 4}
	//res2 := arraySum(arrays2)  // 语法错误，函数无法接收长度超过3的数组


	// 1、切片的定义
	var a[]string
	var b = []int{}
	var c = []bool{}
	var d = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)
	fmt.Println(b == nil)  // 初始化后不为空
	fmt.Println(d == nil)

	// 注意: 切片是引用类型，不支持直接比较，只能和nil比较
	//fmt.Println(a == b)

	// 2、使用 make() 函数构造切片
	/*
	- make([]T, size, cap)
		- T:切片的元素类型
		- size:切片中元素的数量
		- cap:切片的容量

	- 切片的长度与容量
		- 长度: 切片的长度，用 len() 来获取。
		- 容量: 切片的容量，可存放的数组大小，用 cap() 来获取。
	 */
	// slice内部已经分配了10个容量的空间，并且初始化使用了2个
	slice := make([]int, 2, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 总结: 切片的底层其实就是数组，设置容量其实就是数组的长度，
	// 可以先设置好一个可接受的容量，再使用容量内值的个数，相对于数组比较灵活。

	arrays2 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice2 := arrays2[3:5]
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	// 若数组的长度足够时，切片的容量从 low的位置到high的位置
	println(cap(slice2))  // 5

	// 3、切片的赋值拷贝
	s1 := make([]int, 3)
	fmt.Println(s1, "赋值前")

	// 引用传递
	s2 := s1
	s2[0] = 100
	fmt.Println(s2)
	fmt.Println(s1, "赋值后")

	// 4、切片遍历
	s3 := []int{1, 3, 5}
	for _, val := range s3{
		fmt.Println(val)
	}
	for i:=0; i<len(s3); i++{
		fmt.Println(s3[i])
	}

	// 5、append方法为切片追加数据
	s4 := []int{1, 2, 3}
	s4 = append(s4, 4)
	s4 = append(s4, 5)
	num := 0
	TAG:
		for {
			s4 = append(s4, num)
			if num == 20{
				break TAG
			}
			num ++
		}

	fmt.Println(s4)

	// 6、slice底层，自动扩容
	/*
		slice底层原理, 每个切片都会指向一个底层数组，这个数组的容量如果够用就添加元素。
	当底层发数组不能容纳新的元素时，切片就会自动按照一定的策略进行 "扩容", 此时该切片指向的底层数组就
	会更换。"扩容"操作往往是发生在append函数的调用时，所以我们通常都需要用原变量来接收append的返回值。
	*/
	var s5 []int
	for i:=0; i < 10; i++ {
		s5 = append(s5, i)
		// 根据打印所得，每次扩容后都是扩容前的2倍。
		fmt.Printf("%v  len:%d   cap:%d  ptr:%p\n", s5, len(s5), cap(s5), s5)
	}

	// append支持一次性追加多个值
	s6 := []int{1, 2, 3}
	s6 = append(s6, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(s6)

	// 追加切片
	s7 := []int{11, 12, 13, 14}
	s6 = append(s6, s7...)
	fmt.Println(s6)

	// 7、切片的扩容策略
	/*
	- $GOROOT/src/runtime/slice.go 源码
		newcap := old.cap
		doublecap := newcap + newcap
		if cap > doublecap {
			newcap = cap
		} else {
			if old.len < 1024 {
				newcap = doublecap
			} else {
				// Check 0 < newcap to detect overflow
				// and prevent an infinite loop.
				for 0 < newcap && newcap < cap {
					newcap += newcap / 4
				}
				// Set newcap to the requested cap when
				// the newcap calculation overflowed.
				if newcap <= 0 {
					newcap = cap
				}
			}
		}

	- 源码分析
		- 1.首先判断 申请容量 cap 是否大于 2倍 的旧容量old.cap, newcap就是新申请的容量cap。
		- 2.否则判断 旧切片的长度小于 1024, 则容量 newcap 就是 old.cap旧容量的两倍。
		- 3.否则判断 旧切片长度大于等 1024, 则最终容量newcap从旧容量old.cap开始循环增加原来的4分之1。
		- 4.如果最终容量 cap 计算值溢出，则最终容量 cap 就是新申请的容量 cap
	- 注意:
		- 切片扩容机制会根据切片中元素的类型不同去做不同的扩容。比如int和string;
	*/

	// 8、使用copy()函数复制切片
	/// 由于 切片 是引用类型，赋值给一个新的变量时，会实现数据共享，
	/// 这样会导致修改新的变量对原来的变量造成影响，使用copy可以解决这个问题
	s8 := []int{1, 2, 3}
	s9 := s8
	fmt.Println(s8, s9)

	s9[0] = 100
	fmt.Println(s8, s9)

	// 使用copy来实现
	s10 := make([]int, 3, 3)
	// 注意，copy(aSlice, bSlice)  a切片的长度如果比b的小，则复制a中相同长度的数据
	copy(s10, s9)
	fmt.Println(s10, s9)
	s10[0] = 200
	fmt.Println(s10, s9)


	// 9、同一个切片切出来的切片，会切片数据共享
	s11 := []int{1, 2, 3, 4, 5, 6}
	s12 := s11[:3]
	fmt.Println(s11, s12)
	s12[0] = 200
	fmt.Println(s11, s12)

	// 10、从切片中删除元素
	/// go语言中没有专门的删除切片的方法，可以通过切片来实现
	s13 := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 1010}
	/// 删除索引为2的元素
	// 先切片得到只有索引为 0, 1 的切片，然后再跳过索引2，从索引3开始将后面所有的值追加进来
	s13 = append(s13[:2], s13[3:]...)

	// 11、练习题
	/// 11.1 请写出下面代码的输出结果
	var s14 = make([]string, 5, 10)
	for i:=0;i<10;i++{
		s14 = append(s14, fmt.Sprintf("%v", i))
	}

	/// 11.2 请使用内置函数sort对以下数组进行排序;
	var a = [...]int{3, 7, 8, 9, 1}
	a = sort.Sort()
}
