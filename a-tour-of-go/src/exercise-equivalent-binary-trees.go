package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// tree package
// type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	for i := range t {
		fmt.Println(i)
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
}
