package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {
		by3 := division(i, 3)
		by5 := division(i, 5)
		by15 := division(i, 15)

		switch {
		case by15:
			fmt.Println("FizzBuzz")
		case by3:
			fmt.Println("Fizz")
		case by5:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func division(i, j int) bool {
	if i%j == 0 {
		return true
	} else {
		return false
	}
}
