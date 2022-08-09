package main

import "fmt"

func main() {
	f := FV()
	fmt.Println(f(1, 2))
	fmt.Println(f(1, 2))

}
func FV() func(c, d int) (int, int) {
	a := 1
	b := 1
	return func(c, d int) (int, int) {
		c = c + a
		d = d + b
		a = a + 1
		b = b + 2
		fmt.Println("相加结果：", c+d)
		return a, b
	}

}
