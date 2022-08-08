package main

import (
	"fmt"
	"time"
)

func main() {
	go counter()
	fmt.Println("Press the Enter Key to stop anytime")
	fmt.Scanln()
}

func counter() {
	i := 0
	for {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
		i++
	}
}
