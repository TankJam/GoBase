package main

import (
	"encoding/json"
	"fmt"
)

/*
	结构体可见性和json:
		- 重点: Go语言数据转为json
*/

// 定义 学生 结构体
type student struct{
	Id int
	Name string
}

// student构造函数
func initStudent(id int, name string) *student{
	return &student{
		Id:   id,
		Name: name,
	}
}

// 定义 班级 结构体
type class struct{
	Title string `json:"title"`  // json的字段key
	// 数组中存对象
	/// 注意: 反序列化时，json的key必须要与json: "字段名" 一一对应
	Students []student `json:"studentArrays" db:"student" xml:"ss"`
}

// class构造函数
func initClass(title string) *class{
	return &class{
		Title:    title,
		Students: make([]student, 0, 20),  // 创建长度为0，容量为20的空数组
	}
}


func main() {
	// 1.实例化一个班级对象c1
	c1 := initClass("终极一班")

	// 2.往班级c1中添加学生
	for i := 0; i < 10; i++ {
		tmpStu := initStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, *tmpStu)
	}

	fmt.Printf("%#v\n", c1)

	// 3.序列化  json.Marshal -> type -> []uint8 -> bytes
	/// 问题: json序列化对象时，对象的属性必须是首字母大写，否则无法识别
	/// 解决: 定制化json模块
	data, err := json.Marshal(c1)
	if err != nil{
		fmt.Println("json marshal failed, err: ", err)
		return
	}
	fmt.Printf("%T\n", data)  // bytes
	fmt.Printf("%s\n", data)


	// 4.反序列化 json.Unmarshal([]uint8, 结构体指针)
	/// 需要反序列化的json数据
	/// Go语言中json数据用 `{key:value}` 表示
	jsonStr := `{"Title":"火箭101","studentArrays":[{"Id":0,"Name":"stu00"},{"Id":1,"Name":"stu01"},{"Id":2,"Name":"stu02"},{"Id":3,"Name":"stu03"},{"Id":4,"Name":"stu04"},{"Id":5,"Name":"stu05"},{"Id":6,"Name":"stu06"},{"Id":7,"Name":"stu07"},{"Id":8,"Name":"stu08"},{"Id":9,"Name":"stu09"}]}`
	/// 1) 先实例化反序列化后映射的空结构体
	var c2 class // 用于接收反序列化后的结果
	/// 2) 将json数据先转为bytes类型的数据，然后传给c2指针
	err = json.Unmarshal([]byte(jsonStr), &c2)  // &c2 接收者可以将bytes数据传给指针
	if err != nil{
		fmt.Println("json unmarshal failed, err:", err)
		return
	}

	fmt.Printf("%#v\n", c2)
	fmt.Println(c2)

}
