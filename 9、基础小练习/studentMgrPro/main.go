package main

import (
	"fmt"
	"os"
)

/*
	学员管理系统: (结构体包管理 版本)
		- 命令行菜单
		- 添加学生
		- 修改学生
		- 删除学生
		- 展示学生
*/


// showMenu  展示菜单
func showMenu(){
	fmt.Println("~~~~~~学员管理系统欢迎您~~~~~~")
	fmt.Println("1.添加学生")
	fmt.Println("2.修改学生")
	fmt.Println("3.删除学生")
	fmt.Println("4.展示学生")
	fmt.Println("5.退出程序")
}

// userInput 用户输入功能
func userInput() *Student {
	var (
		id int64
		name, class string
		age int8
	)

	fmt.Println("请按照提示录入信息")
	fmt.Print("请输入姓名: ")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄: ")
	fmt.Scanln(&age)
	fmt.Print("请输入学号: ")
	fmt.Scanln(&id)
	fmt.Print("请输入班级: ")
	fmt.Scanln(&class)
	// NewStudent 返回的是指针
	newStudent := NewStudent(name, class, id, age)
	return newStudent
}



func main() {
	var choose string
	stuMgr := NewStudentManager()
	for {
		showMenu()
		// 让用户选择编号
		fmt.Println("请选择功能编号!")
		fmt.Scanln(&choose)
		switch choose {

		case "1":
			// 1.添加学生
			newStu := userInput()
			stuMgr.AddStudent(newStu)

		case "2":
			// 修改学生
			newStu := userInput()
			stuMgr.UpdateStudent(newStu)
		case "3":
			// 删除学生
			newStu := userInput()
			stuMgr.DeleteStudent(newStu)
		case "4":
			// 展示学生
			stuMgr.ShowStudents()
		case "5":
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入!")
		}
	}
	
}
