package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{X: 1, Y: 2})
	v := Vertex{X: 1, Y: 2}
	v.X = 4
	v.Y = 17
	fmt.Println(v.X, v.Y)
}
