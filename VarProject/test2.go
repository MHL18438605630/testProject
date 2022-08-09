package main

import "fmt"

func A() {
	fmt.Println("the name of function is A")
}

func B(a, b int) {
	fmt.Println("the name of function is B , Result = ", a+b)
}

func C(a, b int) int {
	c := a + b
	fmt.Println("the name of function is C , Result = ", c)
	return c
}

func main() {
	var f1 func()
	var f2 func(int, int)
	var f3 func(int, int) int

	f1 = A
	f2 = B
	f3 = C

	f1()
	f2(1, 2)
	fmt.Println(f3(3, 4))

}
