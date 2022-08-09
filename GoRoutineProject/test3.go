package main

import "fmt"

type A interface {
	Test(i int)
}

type B struct {
}

type C struct {
}

func (b *B) Test(i int) {
	fmt.Println("B 中 i = ", i)
}

func (c *C) Test(i int) {
	fmt.Println("C 中 i = ", i)
}

func main() {
	ch := make(chan func(int), 2)

	b := new(B)
	c := new(C)

	ch <- b.Test
	ch <- c.Test

	close(ch)

	for a := range ch {
		a(1)
	}

}
