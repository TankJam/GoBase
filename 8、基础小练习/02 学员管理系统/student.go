package main

import "fmt"

/*
	学生结构体 与 构造方法
*/

// Student 学生结构体
type Student struct {
	id int64
	name, class string
	age int8
}

// NewStudent 学生构造函数
func NewStudent(name, class string, id int64, age int8) *Student{
	return &Student{  // 返回Student地址
		id:    id,
		name:  name,
		class: class,
		age:   age,
	}
}



/*
	通过一个结构体，统一管理以下所有方法
*/

// StudentManager 学生管理结构体
type StudentManager struct {

	// 用户存放所有学生的切片
	allStudents []*Student
}

// NewStudentManager 构造方法
func NewStudentManager() *StudentManager{
	return &StudentManager{
		allStudents: make([]*Student, 0, 100), // 初始化指针数组
	}
}

// AddStudent  添加学生
func (s *StudentManager)AddStudent(stu *Student){
	for _, v := range s.allStudents{
		if v.name == stu.name{
			fmt.Printf("姓名：[%s] 已存在!", v.name)
			return
		}
	}
	s.allStudents = append(s.allStudents, stu)
}

// UpdateStudent  修改学生
func (s *StudentManager) UpdateStudent(stu *Student){
	for index, v := range s.allStudents{
		if v.name == stu.name{
			s.allStudents[index] = stu
		}
	}
	fmt.Printf("姓名为%s的学生不存在\n", stu.name)

}

// DeleteStudent  删除学生
func (s *StudentManager) DeleteStudent(stu *Student){
	for index, v := range s.allStudents{
		if v.name == stu.name{
			// 1.先获取 学生切片里当前删除学生前的所有学生，生成新的切片
			/// s.allStudents[:index]
			// 2.再将删除的学生之后的学生追加进去，实现删除切片中学生的效果
			/// s.allStudents[index + 1:]...  将切片打散，一个个传入append中
			s.allStudents = append(s.allStudents[:index], s.allStudents[index + 1:]...)
			return
		}
	}
	fmt.Printf("姓名为%s的学生不存在\n", stu.name)
}


// ShowStudents  展示学生
func (s *StudentManager) ShowStudents(){
	for _, v := range s.allStudents{
		fmt.Printf("学号: [%d]  -  学生: [%s]  -  年龄: [%d]  -  班级: [%s]\n",
			v.id, v.name, v.age, v.class)
	}
}

