package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z2 := z
		z -= (z*z - x) / (2 * z)
		if z-z2 == 0 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
