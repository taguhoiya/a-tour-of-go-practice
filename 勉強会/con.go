package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	for i := 0; i < 2; i++ {
		if n1, ok := <-ch1; ok {
			// 処理1
			fmt.Println(n1)
		} else if n2, ok := <-ch2; ok {
			// 処理2
			fmt.Println(n2)
		} else {
			// 例外処理
			fmt.Println("neither cannot use")
		}
	}
}
