package main

import (
	"fmt"
)

func main() {
	s := "hello"
	b := []byte(s)
	b[0] = 'c'
	fmt.Println(string(b))

	d := "hello"
	e := "c" + d[1:] // 文字列を変更することはできませんが、スライスは行えます。
	fmt.Printf("%s\n", e)

	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// length: 2, capacity: 5のslice
	slice := Array_a[2:4:7]
	fmt.Printf("%s, %v\n", slice, cap(slice))
}