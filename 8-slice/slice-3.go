package main

import "fmt"

func main() {
	//切片的长度len和容量cap不一样
	//长度表示左指针到右指针之间的距离
	//容量表示左指针到底层数组末尾的距离

	var numbers = make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向切片追加一个元素1
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//向一个容量已满的slice追加元素，切片的容量会扩容最初设置的5
	numbers = append(numbers, 6)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("-------------------")

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	//如果一开始没有设置cap，那么会以最初的切片长度作为cap
	numbers2 = append(numbers2, 4)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

}
