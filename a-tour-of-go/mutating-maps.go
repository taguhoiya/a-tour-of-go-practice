package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)

	m[0] = 42
	fmt.Println("The value:", m[0])

	delete(m, 0)
	fmt.Println("The value:", m[0])

	v, ok := m[0]
	fmt.Println("The value:", v, "Present?", ok)

}
