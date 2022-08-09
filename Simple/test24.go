package main

import (
	"fmt"
	"io"
)

type WriteBook interface {
	WriteBook()
}

type ReadBook interface {
	ReadBook()
}

type Book struct {
	Name string
}

func (this *Book) WriteBook() {
	fmt.Println("write a book")
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func main() {
	book1 := Book{"awm"}

	book1.WriteBook()
	book1.ReadBook()

	var w io.Writer

	w.Write([]byte("jjjj"))

}
