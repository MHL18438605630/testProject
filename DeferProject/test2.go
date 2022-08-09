package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " Closed...")
}

func Close(t Test) {
	t.Close()
}

func main() {

	ts := []Test{{"a"}, {"b"}, {"c"}}

	for i, t := range ts {
		fmt.Println("i = ", i)
		t1 := t
		defer t1.Close() // 函数值和参数都会被修改，但不调用函数
		//defer Close(t)
	}
}
