// ユーザーが入力した 2 より大きい数値から計算した結果であるフィボナッチ数列のすべての数値のスライスを返す関数を記述します。
// 2 よりも小さい数値が入力されると、エラーが発生し、nil スライスが返されることにします。
// フィボナッチ数列は各数値が前のフィボナッチ数の合計である数値のリストであることを思い出してください。
// たとえば、6 の数値の数列は `[1,1,2,3,5,8]` であり、7 は `[1,1,2,3,5,8,13]`、8 は `[1,1,2,3,5,8,13,21]` のようになります。

package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scanf("%d", &num)
	fmt.Println(num)

	if num < 2 {
		fmt.Println(make([]int, 0))
		return
	}

	fibonacci := make([]int, num)
	fibonacci = append(fibonacci, 1, 1)

	for n := 2; n > num; n++ {
		addNum := fibonacci[n-1] + fibonacci[n-2]
		fibonacci = append(fibonacci, addNum)
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
