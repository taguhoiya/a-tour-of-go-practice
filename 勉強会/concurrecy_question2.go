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

var quit = make(chan bool)

func fib(ch chan<- int) {
	x, y := 1, 1

	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- x
}

func main() {
	start := time.Now()
	var input string

	ch1 := make(chan int)

	go fib(ch1)

	for {
		num := <-ch1
		fmt.Println(num)
		fmt.Scanf("%s", &input)
		if input == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
