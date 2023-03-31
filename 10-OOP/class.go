package main

import "fmt"

//如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果类的属性/方法首字母大写，表示该属性对外是public，否则是private
	Name  string
	Ad    int
	Level int
}

func (this *Hero) Show() {
	fmt.Println("hero = ", this)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100, Level: 1}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
