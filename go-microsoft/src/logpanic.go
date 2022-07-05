package main

import (
	"fmt"
	"log"
)

func main() {
	log.Panic("Hey, I'm an error log!")
	fmt.Print("Can you see me?")
}
