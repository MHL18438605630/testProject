package main

import "fmt"

func main() {
	TestMap := make(map[int]string)

	TestMap[1] = "A"
	TestMap[2] = "B"
	TestMap[3] = "C"
	TestMap[4] = "D"
	TestMap[5] = "E"

	value, ok := TestMap[6]

	fmt.Println("value =", value)
	fmt.Println("OK =", ok)
}
