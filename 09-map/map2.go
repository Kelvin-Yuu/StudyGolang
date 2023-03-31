package main

import "fmt"

func printMap(cityMap map[string]string) {
	//这里传的是指针
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func main() {

	cityMap := map[string]string{
		"China": "Beijing",
		"Japan": "Tokyo",
		"USA":   "NewYork",
	}
	//遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
	//删除
	delete(cityMap, "Japan")

	//修改
	cityMap["USA"] = "华盛顿"

	fmt.Println("-------------------")

	printMap(cityMap)
}
