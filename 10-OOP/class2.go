package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法和属性

	level int
}

//重定义父类方法的Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Show() {
	fmt.Println("SuperMan name = ", this.name)
	fmt.Println("SuperMan sex = ", this.sex)
	fmt.Println("SuperMan level = ", this.level)
}

func main() {
	h := Human{"zhang3", "female"}

	h.Eat()
	h.Walk()
	fmt.Println("-----------------------------------------")

	// s := SuperMan{Human{"li4", "male"}, 88}
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 88

	s.Walk()
	s.Eat()
	s.Fly()

	s.Show()
}
