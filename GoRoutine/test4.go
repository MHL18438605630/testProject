package main

import "fmt"

func SubGoRoutine(chanVal chan int) {
	chanVal <- 1
	chanVal <- 2
	chanVal <- 3
	chanVal <- 4

}

func main() {

	chanVal := make(chan int, 3)

	go SubGoRoutine(chanVal)
	fmt.Println(<-chanVal)
}
