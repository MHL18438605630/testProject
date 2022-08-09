package main

import "fmt"

type People struct {
	name string
	age  int
	sex  string
}

func main() {
	people := People{"mhl", 24, "girl"}
	people1 := People{"mhl", 24, "girl"}

	fmt.Printf("%p %p\n", &people, &people1)
	fmt.Println(people1 == people) // true

	people3 := new(People)
	people3 = &People{"MHL", 24, "girl"}
	people4 := &People{"MHL", 24, "girl"}

	fmt.Println(people4 == people3) // false

}
