package main

import "fmt"

func AddUpper() func(int) int {
	a := 1
	b := 10
	return func(x int) int {
		c := a + b + x
		a = a + 1
		b = b - 1
		return c
	}
}
func main() {
	f := AddUpper()
	fmt.Println("闭包返回", f(1)) // 21
	fmt.Println("闭包返回", f(2)) // 23
	fmt.Println("闭包返回", f(3)) // 26
}
