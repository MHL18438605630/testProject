package main

import (
	"fmt"
	"time"
)

func SubPrint() {
	fmt.Println(" hello go")
}
func main() {

	go SubPrint()
	fmt.Println("Hello world")
	time.Sleep(time.Second)
}
