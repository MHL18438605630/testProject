package main

import "fmt"

func main() {

	for {
		for i := 0; i < 10; i++ {

			if i == 6 {
				fmt.Println("i = ", i)
				return
			}

			fmt.Println("Happy ....")

		}
	}

}
