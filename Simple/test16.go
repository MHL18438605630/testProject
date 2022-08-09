package main

import "fmt"

func foo(a, b int) (r1, r2 int) {
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)
	r1 = 100
	r2 = 2233
	return

}

func main() {
	ret1, ret2 := foo(2, 3)
	fmt.Println(ret1)
	fmt.Println(ret2)

}
