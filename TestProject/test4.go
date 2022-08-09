package main

import "fmt"

func main() {
	intNums := []int{1, 2, 3, 4, 5}

	for i, num := range intNums {
		fmt.Println("i = ", i, " , num = ", num)
	}
}
