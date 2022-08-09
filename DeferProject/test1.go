package main

import "fmt"

func main() {
	strNum := [5]int{}

	for i := range strNum {
		fmt.Println("i = ", i)
		defer func() {
			fmt.Println(i)
		}()
	}

}
