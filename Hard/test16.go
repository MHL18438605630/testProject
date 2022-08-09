package main

import "fmt"

func main() {
	DealData(TwoValAdd, 1, 2, 3, 4, 5)

}

type Add func(a, b int) int

func DealData(f Add, param ...int) {
	res := 0
	for a, b := range param {
		fmt.Println(a)
		res = f(res, b)
		fmt.Println(res)
	}
	fmt.Println(res)
}

func TwoValAdd(a, b int) int {
	return a + b
}
