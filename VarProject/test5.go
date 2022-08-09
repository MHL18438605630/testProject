package main

import "fmt"

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T 方法")
}
func (t *T) testP() {
	fmt.Println("如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 *T 方法")
}

func main() {
	t1 := T{1}
	t1.testT()
	t1.testP()

}
