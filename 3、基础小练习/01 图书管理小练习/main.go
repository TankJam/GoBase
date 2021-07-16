package main

import (
	"fmt"
	"os"
)

/*
需求
	使用函数实现一个简单的图书管理系统。
	每本书有书名、作者、价格、上架信息，
	用户可以在控制台添加书籍、修改书籍信息、打印所有的书籍列表。
*/

// 1.定义结构体
type book struct{
	title string
	author string
	price float32
	publish bool
}


// initBook 注释规范
// 2.定义新建书本构造函数
func initBook(title, author string, price float32, publish bool) *book{
	return &book{
		title:   title,
		author:  author,
		price:   price,
		publish: publish,
	}
}

// 3.定义一个存放在book指针的切片, 用于存放所有的书籍对象
var allBooks = make([]*book, 0, 20)


// 4.功能实现
// showMenu 1. 打印菜单
func showMenu(){
	fmt.Println("欢迎登录BMS")
	fmt.Println("1.添加书籍")
	fmt.Println("2.修改书籍信息")
	fmt.Println("3.展示所有书籍")
	fmt.Println("4.退出")
}

// userInput 2. 等待用户输入菜单选项
// 定义一个专门用来获取用户输入的书籍信息的
func userInput()*book{
	var (
		title string
		author string
		price float32
		publish bool
	)
	/// 2.1 获取用户输入
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入书名: ")
	fmt.Scanln(&title)

	fmt.Print("请输入作者: ")
	fmt.Scanln(&author)

	fmt.Print("请输入价格: ")
	fmt.Scanln(&price)

	fmt.Print("是否已上架[true|false]: ")
	fmt.Scanln(&publish)

	return initBook(title, author, price, publish)
}

// addBook 3. 添加书籍的函数
func addBook(){
	// 新建书本对象
	book := userInput()

	// 判断书籍是否存在，不存在则添加到书籍切片中
	for _, b := range allBooks{
		if b.title == book.title{
			fmt.Printf("《%s》这本书已经存在了！", book.title)
			return
		}
	}

	allBooks = append(allBooks, book)
	fmt.Printf("添加书籍【%s】成功!\n", book.title)

}

// updateBook 4. 修改书籍的函数
func updateBook(){
	// 获取用户输入的书籍信息
	book := userInput()

	// 循环遍历书籍数组，判断 书籍是否存在，若存在则修改，否则返回书籍不存在
	for index, b := range allBooks{
		if b.title == book.title{
			allBooks[index] = book
			fmt.Printf("书籍【%s】更新成功!", book.title)
			return
		}
	}
	fmt.Printf("书籍【%s】不存在!", book.title)
}


// showBooks 5. 展示书籍列表
func showBooks(){
	var flag string

	if len(allBooks) == 0{
		fmt.Println("没有书籍~~~")
		fmt.Println("是否创建书籍: 输入 [y 是 | f 否] ：")
		fmt.Scanln(&flag)
		if flag == "y"{
			addBook()
		}else if flag == "n"{
			return
		}else{
			return
		}
	}

	for _, b := range allBooks{
		fmt.Printf("书籍:【%s】  作者:【%s】  价格: 【%.2f】  是否上架销售: 【%t】\n",
			b.title, b.author, b.price, b.publish)
	}

}


// 6. 退出 os.Exit(0)

func main() {

	for {
		// 1) 循环展示书籍信息
		showMenu()

		// 2) 让用户选择数字执行对应的功能
		var option int
		fmt.Scanln(&option)
		switch option{
		case 1:
			addBook()
		case 2:
			updateBook()
		case 3:
			showBooks()
		case 4:
			os.Exit(0)  // 0代表关闭程序
		}

	}
	
}
