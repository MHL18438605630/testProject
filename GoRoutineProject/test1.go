package main

import "fmt"

func main() {

	goRoutineChanVal := make(chan string)

	goRoutineChanVal <- "hello world"

	fmt.Println(<-goRoutineChanVal)

}
