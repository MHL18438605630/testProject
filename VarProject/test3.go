package main

import "fmt"

var Args struct {
	a string
	b int
}

func main() {

	args := struct {
		id   int
		name string
		sex  string
	}{
		1,
		"mhl",
		"woman",
	}

	fmt.Println(args.id)
	fmt.Println(args.name)
	fmt.Println(args.sex)

}
