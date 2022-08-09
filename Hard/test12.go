package main

import "fmt"

func main() {
	GetAddSubValue(SubValue)(5, 3)

}

func SubValue(a, b int) int {
	return b - a
}

func GetAddSubValue(f func(a int, b int) int) func(a int, b int) int {
	return func(a, b int) int {
		SubValue := f(a, b)
		AddValue := a + b
		fmt.Println("相加的数据 a+b=", AddValue, ",相减的数据 b-a =", SubValue)
		return AddValue
	}

}
