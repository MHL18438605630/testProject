package main

import "fmt"

func main() {

	stringChanVal := make(chan string)

	go func() {
		stringChanVal <- "Hello World"
		stringChanVal <- "Hello Go"
		close(stringChanVal)
	}()

	for stringVal := range stringChanVal {
		fmt.Println("StringVal = ", stringVal)
	}

}
