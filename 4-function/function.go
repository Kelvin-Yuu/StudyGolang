package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("----  foo1 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

//返回多个返回值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("----  foo2 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c, 777
}

//返回多个返回值，有形参名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----  foo3 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----  foo3 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("Alice", 100)

	fmt.Println("c = ", c)

	ret1, ret2 := foo2("Alice", 100)

	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	ret3, ret4 := foo3("Alice", 100)

	fmt.Println("ret3 = ", ret3, "ret4 = ", ret4)

	ret5, ret6 := foo4("Alice", 100)

	fmt.Println("ret5 = ", ret5, "ret6 = ", ret6)
}
