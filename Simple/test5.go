package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person1 := Person{"MHL", 24}
	fmt.Println(person1)

	var person2 Person
	person2.name = "karma"
	person2.age = 24

	fmt.Println(person2)

	var person3 = Person{"anoxia", 24}
	fmt.Println(person3)

}
