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
	Students []student `json:"student_array" db:"student" xml:"ss"`
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

}
