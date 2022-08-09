package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("A中的i值为: ", i)
		}()
	}

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("B中的i值为: ", i)
		}()
	}

	fmt.Println("主协程")

	wg.Wait()

}
