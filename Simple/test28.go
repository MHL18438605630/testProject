package main

import (
	"fmt"
)

func main() {

	fmt.Println("main函数")
	c := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(c)

	f := func(a, b int) int {
		return a - b
	}

	h := f(5, 666)

	fmt.Println(h)

}
