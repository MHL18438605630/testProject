package main

import "fmt"

func main() {
	var str []string
	str = append(str, "h")
	str = append(str, "e")
	str = append(str, "l")
	str = append(str, "l")
	str = append(str, "o")
	str = append(str, "!!!")

	var in []int
	in = append(in, 1) // 0
	in = append(in, 2) // 1
	in = append(in, 3) // 2
	in = append(in, 4) // 3
	in = append(in, 5) // 4

	fmt.Println(in[1:3])

	fmt.Println(str)

}
