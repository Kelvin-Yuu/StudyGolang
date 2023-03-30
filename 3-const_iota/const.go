package main

import "fmt"

//const 来定义枚举类型
const (
	//可以再const() 添加一个关键字iota，每行的iota会累加1，第一行默认0
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

func main() {
	const length int = 10
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
}
