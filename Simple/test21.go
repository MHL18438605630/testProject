package main

import (
	"fmt"
)

type ManInterface interface {
	Eat()
	Sleep()
	Fly()
}

type Woman1 struct {
	name string
}

type Man struct {
	name string
}

func (this *Woman1) Eat() {
	fmt.Println(this.name + "正在吃饭")
}

func (this *Woman1) Sleep() {
	fmt.Println(this.name + "正在睡觉")
}

func (this *Woman1) Fly() {
	fmt.Println(this.name + "正在飞翔")
}

func (this *Man) Eat() {
	fmt.Println(this.name + "正在吃饭")
}

func (this *Man) Sleep() {
	fmt.Println(this.name + "正在睡觉")
}

func (this *Man) Fly() {
	fmt.Println(this.name + "正在飞翔")
}

func Show(manInterface ManInterface) {
	manInterface.Eat()
	manInterface.Sleep()
	manInterface.Fly()

}

func main() {

	woman222 := Woman1{"MHL"}
	woman222.Eat()
	woman222.Sleep()
	woman222.Fly()

	man := Man{"kkkk"}

	man.Eat()
	man.Sleep()
	man.Fly()

	//man1=&man{"jjjjj"}
	//
	Show(&woman222)
	Show(&man)

}
