package main

import "fmt"

func restFunc() <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		for i := 0; i < 5; i++ {
			res <- 1
		}
	}()

	return res
}

func main() {
	ch := restFunc()
	for i := 0; i < 8; i++ {
		fmt.Println(<-ch)
	}
}
