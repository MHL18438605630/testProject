package main

import (
	"fmt"
)

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int", a)
	case string:
		fmt.Println("the type of a is string", a)
	case float64:
		fmt.Println("the type of a is float", a)
	default:
		fmt.Println("unknown type")
	}

}

func main() {

	x := "Hello world"
	y := 4
	z := 0.01
	getType(y)
	getType(x)
	getType(z)
}
