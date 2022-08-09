package main

import "fmt"

type People struct {
	string
	int
}

func (p *People) PrintPeople() {
	fmt.Println("People = { Name = ", p.string, ", Age = ", p.int, " }")
}

func main() {

	p := People{
		"MHL",
		25,
	}

	p.PrintPeople()

	fmt.Println("People = ", p)
	fmt.Println("Name = ", p.string)
	fmt.Println("Age = ", p.int)

}
