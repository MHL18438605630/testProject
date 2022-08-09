package main

import "fmt"

func main() {

	go func() {
		fmt.Println("子协程的功能......")
	}()

	select {}

	fmt.Println("主协程的功能......")

}
