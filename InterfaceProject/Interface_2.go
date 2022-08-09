package main

import "fmt"

type Test1 interface {
	A(i int)
	B()
}

type Test2 interface {
	A(i int)
	D()
}

type TestStr struct {
	StrName string
}

func (this *TestStr) A(i int) {
	fmt.Printf("The Interface of Test%d: A%d\n", i, i)
}

func (this *TestStr) B() {
	fmt.Println("The Interface of Test1: B")
}

func (this *TestStr) D() {
	fmt.Println("The Interface of Test2: D")
}

func InterAcc1(t Test1) {
	t.A(1)
	t.B()
}

func InterAcc2(t Test2) {
	t.A(2)
	t.D()
}

func main() {

	testStr := TestStr{
		StrName: "MHL",
	}

	InterAcc1(&testStr)
	InterAcc2(&testStr)

}
