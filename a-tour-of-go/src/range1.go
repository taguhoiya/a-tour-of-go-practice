package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // 左シフト 1→10→100
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
