package main

import "fmt"

func main() {
	res := map[int]string{
		1: "kaixin",
		2: "mahuili",
		3: "ABABABAC",
	}
	delete(res, 1)
	res[2] = "ccccc"
	res[3] = "ddddd"

	fmt.Println(res)

}
