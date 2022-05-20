package main

import (
	"fmt"
)

func main() {
	var q = []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	s := []struct{
		i int
		b bool
	}{
		{2, false},
		{3, true},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	var ary []int = q[:2]
	var ary2 []int = q[0:]
	var ary3 []int = q[:]
	fmt.Println(ary)
	fmt.Println(ary2)
	fmt.Println(ary3)

}
