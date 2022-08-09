package main

import (
	"fmt"
)

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println("Val = ", val)
	}
}

func main() {
	//names := []string{"stanley", "david", "oscar"}
	//
	//vars := make([]interface{}, len(names))
	//
	//for i, val := range vars {
	//	vars[i] = val
	//}
	//
	//PrintAll(vars)

	//names := []string{"stanley", "david", "oscar"}
	//vals := make([]interface{}, len(names))
	//for i, v := range names {
	//	vals[i] = v
	//}
	//PrintAll(vals)

	names := []string{"MHL", "ABC", "DEF", "GHI"}

	vars := make([]interface{}, len(names))

	for i, name := range names {
		vars[i] = name
	}
	PrintAll(vars)

}
