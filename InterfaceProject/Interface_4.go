package main

import "fmt"

type A struct {
	aName string
}

type AInter interface {
	PrintA()
}

func (a *A) PrintA() {
	fmt.Println("Name = ", a.aName)
}

func main() {

}
