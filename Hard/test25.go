package main

import "fmt"

func main() {
	a := 1
	b := 2
	for i := 0; i < 5; i++ {
		if a == 1 {
			if b == 2 {
				fmt.Println("hello 1 ")
				break
			}
			fmt.Println("hello 2 ")
		}
		fmt.Println("hello 3 ")
	}

	fmt.Println("hello 4 ")

}
