package main

import "fmt"

func main() {
	var funcName interfaceFunc
	funcName = new(structName)

	fmt.Println(funcName.add(3, 5))
	fmt.Println(funcName.sub(5, 3))

}

type interfaceFunc interface {
	sub(a, b int) int
	add(a, b int) int
}

type structName struct {
}

func (structName1 structName) sub(a, b int) int {
	return a - b
}

func (structName1 structName) add(a, b int) int {
	return a + b
}

