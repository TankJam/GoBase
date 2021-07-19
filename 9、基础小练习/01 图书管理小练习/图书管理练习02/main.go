package main

import (
	"fmt"
	"os"
)

type book struct{
	title string
	author string
	price float32
	publish bool
}

func initBook(title, author string, price float32, publish bool) *book{
	return &book{
		title:   title,
		author:  author,
		price:   price,
		publish: publish,
	}
}


// 存放书籍的切片
var allBooks = make([]*book, 0, 20)


// showMenu 查看图书馆菜单
func showMenu(){
	fmt.Println("----- 图书管理系统 -----")
	fmt.Println("1.添加书籍")
	fmt.Println("2.修改书籍")
	fmt.Println("3.查看书籍")
	fmt.Println("4.退出系统")
}

func userInput() *book{
	var (
		title string
		author string
		price float32
		publish bool
	)
	fmt.Println("请输入书籍名称: ")
	fmt.Scanln(&title)
	fmt.Println("请输入作者名称: ")
	fmt.Scanln(&author)
	fmt.Println("请输入书籍价格: ")
	fmt.Scanln(&price)
	fmt.Println("书籍是否已上架: 输入[true|false]")
	fmt.Scanln(&publish)

	book := initBook(title, author, price, publish)
	return book
}

// addBook 添加书籍
func addBook(){
	book := userInput()
	for _, b := range allBooks{
		if b.title == book.title{
			fmt.Printf("书籍[%s]已存在!", book.title)
			return
		}
	}
	allBooks = append(allBooks, book)
	fmt.Printf("添加书籍[%s]成功!", book.title)
}

// 修改书籍
func changeBook(){
	book := userInput()
	for index, b := range allBooks{
		if b.title == book.title{
			allBooks[index] = book
			fmt.Printf("书籍[%s]修改成功!", book.title)
			return
		}
	}
	fmt.Printf("书籍[%s]不存在!", book.title)
}

// 查看书籍
func showBooks(){
	if len(allBooks) == 0{
		fmt.Println("没有书籍~")
		return
	}
	for _, b := range allBooks{
		fmt.Printf("书籍: [%s] - 价格: [%.2f] - 作者: [%s] - 是否已上架: [%t] \n",
			b.title, b.price, b.author, b.publish)
	}
}


func main() {
	for {
		showMenu()
		var choose int
		fmt.Scanln(&choose)
		switch choose{
		case 1:
			addBook()
		case 2:
			changeBook()
		case 3:
			showBooks()
		case 4:
			os.Exit(0)
		}
	}
}
