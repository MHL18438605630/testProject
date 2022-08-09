package main

import (
	"fmt"
)

type Human struct {
	name string
	sex  string
	age  int
}

func (this *Human) Eat() {
	fmt.Println(this.name + "需要吃饭")
}

func (this *Human) Walk() {
	fmt.Println(this.name + " 正在走路")
}

type Woman struct {
	Human
	level int
}

func (this *Woman) Eat() {
	fmt.Println(this.name + " 正在吃饭饭！！！1")
}

func (this *Woman) Fly() {
	fmt.Println(this.name + "正在飞翔")
}
func main() {
	woman := Woman{Human{"MHL", "GIRL", 23}, 1}

	var woman1 Woman
	woman1.name = ""
	woman1.sex = ""
	woman1.age = 12
	woman1.level = 3
	woman.Eat()
	woman.Fly()
	woman.Walk()

}
