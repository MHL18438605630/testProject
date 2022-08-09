package main

import (
	"fmt"
	"os"
)

func do() error {
	f, err := os.Open("./book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}()
	}

	f, err = os.Open("./another-book.txt")

	if err != nil {
		return err
	}
	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}()
	}
	return nil
}

func main() {

	fmt.Println(do())

}
