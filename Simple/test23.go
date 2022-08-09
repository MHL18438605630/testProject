package main

import (
	"fmt"
)

func haiPrint(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "type is int")
		case string:
			fmt.Println(arg, "type is string")
		case int64:
			fmt.Println(arg, "type is int64")
		case float64:
			fmt.Println(arg, "type is float64")
		default:
			fmt.Println(arg, "type is unknown")
		}
	}
}
func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	//Go语言函数可变参数，可以传入任意个数与任意类型
	haiPrint("Hello", "HaiCoder", 3, 100.1)
}
