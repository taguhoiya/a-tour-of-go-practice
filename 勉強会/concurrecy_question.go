// コンカレンシーを実装する改良バージョン。 現時点では、かかる時間は数秒 (15 秒以内)のはずです。 バッファーありのチャネルを使用する必要があります。

// ユーザーが fmt.Scanf() 関数を使用してターミナルに quit と入力するまでフィボナッチ数を計算する新しいバージョンを作成します。
// ユーザーが Enter キーを押した場合は、新しいフィボナッチ数を計算する必要があります。 つまり、1 から 10 までのループはなくなります。
// 2 つのバッファーなしチャネルを使用します。
// 1 つはフィボナッチ数を計算するためのものであり、もう 1 つはユーザーからの "終了" メッセージを待機するためのものです。
// select ステートメントを使用する必要があります。

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64, ch chan<- float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- x
}

func main() {
	start := time.Now()

	ch := make(chan float64, 15)

	for i := 1; i < 15; i++ {
		go fib(float64(i), ch)
	}

	for i := 1; i < 15; i++ {
		n := <-ch
		fmt.Printf("Fib(%v): %v\n", i, n)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
