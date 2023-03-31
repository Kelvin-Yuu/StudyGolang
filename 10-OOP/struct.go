package main

import "fmt"

//声明一种数据类型myint，是int的一个别名
// type myint int

//定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	//传递一个book的副本
	book.auth = "li4"
}

func changeBook2(book *Book) {
	//传递一个book的指针
	book.auth = "li4"
}

func main() {
	// var a myint = 10
	// fmt.Println("a = ", a)
	// fmt.Printf("type of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
