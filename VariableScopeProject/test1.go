package main

import "fmt"

func main() {

	err := "error1"

	if err := "err2"; err == "err2" {
		fmt.Println("err = ", err)
	}

	fmt.Println("err = ", err)

}
