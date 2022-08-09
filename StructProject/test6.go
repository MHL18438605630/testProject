package main

import (
	"flag"
	"fmt"
)

func main() {
	var indexRuneTests = []struct {
		s    string
		rune rune
		out  int
	}{
		{"a A x", 'A', 2},
		{"some_text=some_value", '=', 9},
		{"a", 'a', 3},
		{"ab", 'h', 4},
	}

	fmt.Println(len(indexRuneTests))

	flag.Args()

}
