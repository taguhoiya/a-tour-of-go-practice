package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("no-file.go")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()
}
