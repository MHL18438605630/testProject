package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	f, err1 := os.Create("test.out")

	if err1 != nil {
		panic(err1)
	}

	defer f.Close()

	err1 = trace.Start(f)
	if err1 != nil {
		panic(err1)
	}

	trace.Stop()

	fmt.Println("hello world!")

}
