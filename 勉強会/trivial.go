package main

import (
	"fmt"
	"bufio"
	"os"
)

var (
	scanner *bufio.Scanner
)

func main()  {
	scanner = bufio.NewScanner(os.Stdin)

	sendCount := 5
	fmt.Printf("Please enter %d names.\n", sendCount)
	s := 0
	for i := 0; i < sendCount; i++ {
		scanner.Scan()
		name := scanner.Text()
		if i == 4 {
			return
		} else {
			fmt.Println(name)
			s++
		}
	}
}