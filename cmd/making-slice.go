package main

import (
	"fmt"
)

func main() {
	var a []int = make([]int, 0, 5)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
