package main

import "fmt"

func main() {
	//声明slice1是一个切片，并且初始化
	slice1 := []int{1, 2, 3}

	//声明slice2是一个切片，但没有给slice分配空间
	var slice2 []int
	slice2 = make([]int, 3) //开辟三个空间，默认值为0

	//声明slice3是一个切片，同时分配三个空间，默认值0
	var slice3 []int = make([]int, 3)
	slice4 := make([]int, 3)

	fmt.Printf("slice1 len = %d, slice = %v\n", len(slice1), slice1)
	fmt.Printf("slice2 len = %d, slice = %v\n", len(slice2), slice2)
	fmt.Printf("slice3 len = %d, slice = %v\n", len(slice3), slice3)
	fmt.Printf("slice4 len = %d, slice = %v\n", len(slice4), slice4)

	//判断一个slice是否为空
	var slice5 []int

	if slice5 == nil {
		fmt.Println("slice5是一个空切片")
	} else {
		fmt.Println("slice5是有空间的")
	}
}
