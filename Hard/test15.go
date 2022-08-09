package main

import "fmt"

func main() {

	a := do(wrapper("+"), 1, 2, 3, 4)
	fmt.Println(a)

}

type callTwoInt func(i, j int) int

func do(f callTwoInt, param ...int) int {
	res := 0
	for _, val := range param {
		res = f(res, val)
	}
	return res

}
func operate(i, j int, op string) int {
	if op == "+" {
		return i + j
	} else if op == "-" {
		return i - j
	}
	return 0
}

func wrapper(op string) func(i, j int) int {
	return func(i, j int) int {
		return operate(i, j, op)
	}

}