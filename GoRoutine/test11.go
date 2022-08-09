package main

import "fmt"

func AB(a, b int) {
	c := a + b
	fmt.Println(fmt.Sprintf("%d + %d = %d", a, b, c))
	panic("happened is error")
}

func main() {

	done := make(chan bool)

	go func(a, b int) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("恢复.. ", err)
			}
			done <- true
		}()
		AB(a, b)
	}(2, 3)

	<-done

}
