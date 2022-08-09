package main

import "fmt"

func Judge(c1, c2, c3, quit chan int) {
	i, j, k := 1, 2, 3
	for {
		select {
		case c1 <- i: // 成功写入数据
			fmt.Printf("C1写入成功")
		case c2 <- j: // 成功写入数据
			fmt.Printf("C2写入成功")
		case c3 <- k: // 成功写入数据
			fmt.Printf("C3写入成功")
		case <-quit: // 成功读取数据
			fmt.Printf("quit读取成功")
			return
		}

	}

}

func main() {

	c1 := make(chan int, 2)
	c2 := make(chan int, 2)
	c3 := make(chan int, 2)
	quit := make(chan int, 2)

	go func() {
		for i := 0; i < 1; i++ {
			fmt.Println("读取c1", <-c1)
			fmt.Println("读取c2", <-c2)
			fmt.Println("读取c3", <-c3)
		}

		quit <- 0
	}()

	Judge(c1, c2, c3, quit)

}
