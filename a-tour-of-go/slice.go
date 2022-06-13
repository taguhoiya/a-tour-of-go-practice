// 配列は固定長。スライスは可変長
// スライスは配列よりも一般的

package main

import (
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:2]
	fmt.Println(s)
}

// slices are like references to arrays
// スライスは配列への参照なので、スライスを変更すると配列も全て変更される
