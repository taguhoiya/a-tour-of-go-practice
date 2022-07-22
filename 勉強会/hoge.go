package main

import "fmt"

func main() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		switch {
		case val < 0:
			panic("Invalid number")
		case val == 0:
			fmt.Printf("%s\n", "0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}

// ユーザーに数を尋ねるプログラムを作成してください。下のコードを修正して作成してみてください
// 	繰り返し整数の入力を求めます。 ループの終了条件は、ユーザーが負の数を入力した場合です。
//  ユーザーが負の数を入力したら、プログラムをクラッシュさせます。 その後、スタック トレース エラーを出力します。
//  数が 0 の場合は、`0 is neither negative nor positive `と出力し、 数の要求を続けます。
//  数が正の値の場合は、`You entered: X` と出力し、 数の要求を続けます。(X は入力された数)
