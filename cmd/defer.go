package main

import "fmt"

func main() {
	defer fmt.Println("it's deferred")

	fmt.Println("hello")
}

// deferはLIFOの法則でスタックされる。
