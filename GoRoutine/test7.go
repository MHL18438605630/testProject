package main

import (
	"fmt"
	"sync"
)

func SubRoutine(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("routine %d:%d\n", x, i)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go SubRoutine(1, &wg)
	go SubRoutine(2, &wg)

	wg.Wait()
	for i := 0; i < 5; i++ {
		fmt.Println("main routine...", i)
	}

}
