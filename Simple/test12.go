package main

import "fmt"

func main() {

	i := 3

	switch i {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 4:
		fmt.Println("D")
	case 5:
		fmt.Println("E")
	default:
		fmt.Println("default")
	}

}
