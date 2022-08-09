package main

func main() {

	ch1 := make(chan []int)

	a := []int{1, 2, 3, 4}

	ch1 <- a

}
