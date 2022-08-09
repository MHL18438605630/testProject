package main

import "fmt"

func main() {

	var mapVal map[string]string
	mapVal = make(map[string]string)

	mapVal["china"] = "beijing"

	fmt.Println(mapVal)

	var mapVal1 map[string]string = map[string]string{}
	mapVal1["France"] = "Italy"
	fmt.Println(mapVal1)

	mapVal2 := make(map[string]string)
	mapVal2["hhh"] = "jjj"

	mapVal3 := map[string]string{
		"AAA": "bbbb",
	}
	fmt.Println(mapVal3)

}
