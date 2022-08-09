package main

import "fmt"

func AddNums(args ...int) {
	sum := 0
	for _, val := range args {
		sum += val
	}
	fmt.Println("sum = ", sum)
}

func main() {
	AddNums(1, 2, 3, 4, 5, 6)

}
