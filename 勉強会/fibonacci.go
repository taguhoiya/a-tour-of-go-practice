// この最初の課題では、数値からフィボナッチ数列を計算するプログラムを記述します。
// これは、新しい言語を学習する際に一般的なプログラミングの練習として行われるコーディングです。
// ユーザーが入力した 2 より大きい数値から計算した結果であるフィボナッチ数列のすべての数値のスライスを返す関数を記述します。
// 2 よりも小さい数値が入力されると、エラーが発生し、nil スライスが返されることにします。
// フィボナッチ数列は各数値が前のフィボナッチ数の合計である数値のリストであることを思い出してください。
// たとえば、6 の数値の数列は 1,1,2,3,5,8 であり、7 は 1,1,2,3,5,8,13、8 は 1,1,2,3,5,8,13,21 のようになります。

package main

import (
	"fmt"
)

func main() {
	val := 0
	var answer []int
	fmt.Scanf("%d", &val)
	var (
		num     int = 1
		nextNum int
	)
	if val < 2 {
		fmt.Println(answer)
	} else {
		for i := 0; i <= val; i++ {
			num, nextNum = nextNum, num+nextNum
			answer = append(answer, nextNum)
		}
		fmt.Println(answer)
	}
}

// package main

// import "fmt"

// func fibonacci(n int) []int {
// 	if n < 2 {
// 		return make([]int, 0)
// 	}

// 	nums := make([]int, n)
// 	nums[0], nums[1] = 1, 1

// 	for i := 2; i < n; i++ {
// 		nums[i] = nums[i-1] + nums[i-2]
// 	}

// 	return nums
// }

// func main() {
// 	var num int

// 	fmt.Print("What's the Fibonacci sequence you want? ")
// 	fmt.Scanln(&num)
// 	fmt.Println("The Fibonacci sequence is:", fibonacci(num))
// }
