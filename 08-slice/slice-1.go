package main

import "fmt"

func printArray(myArray []int) {

	//切片形参是引用传递！！！
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

func main() {
	//动态数组， 切片 slice
	myArray := []int{1, 2, 3, 4}

	fmt.Printf("myArray type is %T\n", myArray)
	printArray(myArray)

	fmt.Println("========")
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}
