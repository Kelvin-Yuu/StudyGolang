package main

import "fmt"

func main() {
	var a string

	//pair<statictype:string, value:"aceld">
	a = "aceld"

	//pair<statictype:string, value:"aceld">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
