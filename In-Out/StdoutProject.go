package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	argValues := os.Args

	if len(argValues) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = argValues[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
