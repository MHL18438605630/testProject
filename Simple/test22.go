package main

import "fmt"

func GetValue(arg interface{}) {

	value, ok := arg.(string)
	if !ok {
		fmt.Println("类型不匹配")
	} else {
		fmt.Println("value =", value)
	}
}

func main() {
	GetValue("MHL")

}
