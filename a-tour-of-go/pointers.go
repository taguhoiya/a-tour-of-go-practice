package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	var p *int = &i // point to i: pはポインタ変数
	fmt.Println(*p) // read i through the pointer: 間接参照

	// 	ポインタ変数が指し示すメモリアドレスは先頭のメモリアドレスだけです。　例： &i = "0x1400012c008"
	// 	そのメモリアドレスから何バイトのデータがあるかは型によって異なるため、何型かの情報がなければ、そのデータにアクセスすることができなくなってしまいます。
	// したがって、このようなことを避けるため、ポインタ変数には明示的にその指し示す型を指定する必要があるのです。

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
