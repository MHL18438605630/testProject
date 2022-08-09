package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello world"
	Wrapper(sayHello)(s)

}

// 中间件

func sayHello(s string) {
	fmt.Println(s)
}

func Wrapper(f func(string)) func(string) {
	return func(s string) {
		a := strings.ToUpper(s)
		f(a)
		fmt.Println("middleware executed")
	}
}
