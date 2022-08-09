package main

import "fmt"

func main() {

	ChanVal := make(chan int, 4)

	go func() {
		for i := 0; i < 10; i++ {
			ChanVal <- i
			fmt.Printf("i = %d \n", i)
			//close(ChanVal)
		}

		close(ChanVal)
	}()

	for {
		if date, ok := <-ChanVal; ok {
			fmt.Println("date = ", date)
		} else {
			break
		}
	}

	fmt.Println("Main Finish ....")

}
