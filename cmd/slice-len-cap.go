package main

import ("fmt")

func main() {
	s:= []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[2:3] // s[2:3]が元となる配列
	printSlice(s)

	s = s[:4]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// スライスは長さ( length )と容量( capacity )の両方を持っています。
// スライスの長さは、それに含まれる要素の数です。
// スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素数です。
// スライスでコロンの後ろがcapacity超えてたらエラーになる
// [1:4] => コロンの前はcapacity決める、このスライスのlength: 3、capacity: 4
// 空スライスはnilで[]を表す