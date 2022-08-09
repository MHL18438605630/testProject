package main

import "fmt"

func Add(args ...interface{}) {
	fmt.Println(" args = ", args)
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Println("V = ", v)
			sum := 0
			for _, val := range args {
				sum += val.(int)
			}
			fmt.Println("sum = ", sum)
		case string:
			fmt.Println("V = ", v)
			fmt.Println("args = ", arg)
		default:
			fmt.Println("V = ", v)
		}
	}

}

func ADD(args ...int) {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	fmt.Println("sum = ", sum)
}

func main() {
	a := []int{1, 2, 3}
	Add(a)

	ADD(1, 2, 3, 4, 5)

}
