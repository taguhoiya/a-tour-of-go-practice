// この最初の課題では、数値からフィボナッチ数列を計算するプログラムを記述します。
// これは、新しい言語を学習する際に一般的なプログラミングの練習として行われるコーディングです。
// ユーザーが入力した 2 より大きい数値から計算した結果であるフィボナッチ数列のすべての数値のスライスを返す関数を記述します。
// 2 よりも小さい数値が入力されると、エラーが発生し、nil スライスが返されることにします。
// フィボナッチ数列は各数値が前のフィボナッチ数の合計である数値のリストであることを思い出してください。
// たとえば、6 の数値の数列は 1,1,2,3,5,8 であり、7 は 1,1,2,3,5,8,13、8 は 1,1,2,3,5,8,13,21 のようになります。

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
		fmt.Println(answer)
	}
}