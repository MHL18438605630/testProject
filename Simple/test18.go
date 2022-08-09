package main

import "fmt"

func main() {
	a := 10
	b := 20
	swap1(a, b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)

	var slien []int = make([]int, 3)
	if slien == nil {

	}
}
func swap1(a, b int) {
	var tmp int
	tmp = a
	a = b
	b = tmp
}
func swap2(a, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}
