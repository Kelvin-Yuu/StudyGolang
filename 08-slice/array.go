package main

import "fmt"

func printArray(myArray [10]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
}

func main() {
	//固定长度数组
	var myArray1 [10]int

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	myArray2 := [10]int{1, 2, 3, 4}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//查看数组的数据类型

	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)

	printArray(myArray2)
}
