package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum() {
	fmt.Println("...")

	// 占いにかかる時間はランダム
	rand.Seed(time.Now().Unix())
	// ランダムな時間スリープさせる
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	fmt.Printf("Today's your lucky number is %d!\n", num)
}

func main() {
	fmt.Println("what is today's lucky number?")
	go getLuckyNum()

	time.Sleep(time.Second * 5)
}
