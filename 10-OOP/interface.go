package main

import "fmt"

//本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类
type Cat struct {
	color string
}

//实现接口方式
func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	color string
}

//实现接口方式
func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("Color = ", animal.GetColor())
	fmt.Println("Kind = ", animal.GetType())
}

func main() {
	// var animal AnimalIF //接口的数据类型，父类指针

	// animal = &Cat{"Yellow"}
	// animal.Sleep()

	// animal = &Dog{"Black"}
	// animal.Sleep()

	cat := Cat{"Yellow"}
	dog := Dog{"Black"}
	showAnimal(&cat)
	showAnimal(&dog)
}
