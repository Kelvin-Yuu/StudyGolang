package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s1 := s[0:2]

	fmt.Println("s:", s, " s1:", s1)
	//以上方法截取slice，s和s1实际上是指向同一块内存（浅拷贝
	s1[0] = 100
	fmt.Println("s:", s, " s1:", s1)

	//copy() 可以将底层数组的slice一起进行拷贝（深拷贝
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println("s2:", s2)

	s2[0] = 1
	fmt.Println("s2:", s2)

}
