package main

import (
	"fmt"
	"sync"
)

func A(i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("恢复 \n", err)
		}
	}()
	fmt.Println("I am A ", i)
	panic("崩溃")
}

func main() {

	var wg sync.WaitGroup
	fmt.Println("It is main goroutine ....")

	wg.Add(1)
	go func(i int) {
		defer wg.Done()
		A(i)
	}(1)
	wg.Wait()
	fmt.Println("执行完了")

}
