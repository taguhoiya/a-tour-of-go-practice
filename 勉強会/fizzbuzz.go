// まず、1 から 100 までの数を出力するプログラムを作成し、次のように変更します。
// 数が 3 で割り切れる場合は Fizz を出力します。
// 数が 5 で割り切れる場合は Buzz を出力します。
// 数が 3 と 5 の両方で割り切れる場合は FizzBuzz を出力します。
// 前のどのケースとも一致しない場合は、数を出力します。

package main

import (
	"fmt"
)

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		switch {
// 		case i%3 == 0 && i%5 == 0:
// 			fmt.Println("FizzBuzz")
// 		case i%3 == 0:
// 			fmt.Println("Fizz")
// 		case i%5 == 0:
// 			fmt.Println("Buzz")
// 		default:
// 			fmt.Println(i)
// 		}
// 	}
// }

func main() {
	for i := 1; i <= 100; i++ {
		by3 := division(i, 3)
		by5 := division(i, 5)
		by15 := division(i, 15)

		switch {
		case by15:
			fmt.Println("FizzBuzz")
		case by3:
			fmt.Println("Fizz")
		case by5:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func division(i, j int) bool {
	if i%j == 0 {
		return true
	} else {
		return false
	}
}
