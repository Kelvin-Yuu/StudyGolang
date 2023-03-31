package main

import "fmt"

//四种变量的声明方式

//声明全局变量  方法一、二、三皆可，方法四只能用于函数体内声明

func main() {
	//方法一：声明一个变量，默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	//方法三：在初始化的时候，可以省去数据类型，通过值自动匹配
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	//方法四：省去var，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)

	var kk, ll = 100, "Alice"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)

}
