package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

/*
	make(map[KeyType]ValueType, [cap])
	make(map[key的类型]value的类型, 容量)
*/

// 1.基本使用
func main() {
	// 1）先定义，后初始化
	scoreMap := make(map[string]int, 8)
	scoreMap["tank"] = 90
	scoreMap["sean"] = 100
	scoreMap["alex"] = 80
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["tank"])
	fmt.Printf("type of a: %T\n", scoreMap)

	// 2）定义并初始化
	userInfo := map[string]string{
		"username": "tank",
		"user": "tank",
		"password": "tank123",
	}
	fmt.Println(userInfo)

	// 3）判断某个key是否存在
	//v, ok := userInfo["username"]
	v, ok := userInfo["username2"]
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("查无此人!")
	}

	// 4）map遍历
	for k, v := range scoreMap{
		fmt.Println(k, v)
	}

	// 5）只想要遍历key的时候
	for k2 := range scoreMap{
		fmt.Println(k2)
	}

	// 6）delete() 通过key删除键值对
	fmt.Println(userInfo)
	delete(userInfo, "user")
	fmt.Println(userInfo)

	// 7）按照指定顺序遍历
	rand.Seed(time.Now().UnixNano())  // 初始化随机数种子
	scoreMap2 := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)  // 生成stu开头的字符串
		value := rand.Intn(100)  // 生成1~99随机的整数
		scoreMap2[key] = value
	}
	fmt.Println(scoreMap2)

	// 取出map中的所有key存入切片keys中
	keys := make([]string, 0, 200)
	for key := range scoreMap2{
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Strings(keys)

	// 按照排序后的key遍历map
	for _, key := range keys{
		fmt.Println(key, scoreMap2[key])
	}

	// 8）元素为map类型的切片
	mapSlice := make([]map[string]string, 3)
	for index, value := range mapSlice{
		fmt.Println(index, value)
		fmt.Printf("index: %d -- value:%v\n", index, value)
	}

	// 对切片中的map元素进行初始化
	//mapSlice[0] = make(map[string]string, 10)
	//mapSlice[0]["name"] = "tank"
	//mapSlice[0]["password"] = "tank9527"
	//mapSlice[0]["address"] = "广东省广州市"
	for index, _ := range mapSlice{
		mapSlice[index] = make(map[string]string, 10)
		mapSlice[index]["name"] = "tank"
		mapSlice[index]["password"] = "tank9527"
		mapSlice[index]["address"] = "广东省广州市"
	}

	for index, value := range mapSlice{
		fmt.Printf("index: %d  --- value: %v\n", index, value)
	}

	// 9) 值为切片类型的map
	sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)

	key := "China"
	value, ok := sliceMap[key]
	// 如果China key不存在，则初始化一个切片
	if !ok{
		value = make([]string, 0, 2)
	}
	value = append(value, "BEIJING", "SHANGHAI", "GUANGZHOU")
	sliceMap[key] = value
	fmt.Println(sliceMap)

	// 10）练习题
	// - 10.1 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	str := "how do you do"
	strSlice := strings.Split(str, " ")
	strMap := make(map[string]int, 10)

	for _, s := range strSlice{
		_, ok := strMap[s]
		if !ok{
			strMap[s] = 1
		}else{
			strMap[s] += 1
		}
	}
	fmt.Println(strMap)

	// - 观察下面代码，写出最终的打印结果。
	// 定义一个map类型的Map类型
	type Map map[string][]int
	// 初始化Map类型
	m := make(Map)
	// 定义并初始化切片
	s := []int{1, 2}
	// 给切片添加新的值
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	// 因为通过 [:1] 切片出来的切片会开启一个新的内存空间，将 s 切片第三个值以及后面的值，添加到新的切片中
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])


}


