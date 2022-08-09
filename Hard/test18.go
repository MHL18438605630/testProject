package main

import "fmt"

func main() {
	//c1 := f(0)
	//c2 := f(0)

	//fmt.Println(f()(0))

	c1 := f(1)
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	//fmt.Println(f()(1))
	//fmt.Println(f()(2))
	//fmt.Println(f()(3))

}

func f(i int) func() int {

	return func() int {
		i++

		fmt.Println("i = ", i)
		return i
	}

}
