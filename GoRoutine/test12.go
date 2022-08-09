package main

import "fmt"

func main() {
	done := make(chan bool)

	go func(a, b int) {
		c := a + b
		fmt.Println(fmt.Sprintf("%d + %d = %d", a, b, c))
		done <- true
	}(2, 3)
	<-done
}

