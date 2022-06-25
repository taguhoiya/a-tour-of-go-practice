package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// しかし、Go言語の場合はその必要はありません。たとえ呼び出し元の変数がポインタ型じゃなくても、レシーバがポインタ型である場合は Go言語が勝手に判断して、
// それをポインタ型として解釈してくれます。（もちろん上記のように、呼び出し元の変数を律儀にポインタ型に書き換えたとしても問題はありません。）

// 呼び出し元の変数がポインタ型でなく、値型であってもレシーバがポインタ型ならgo言語はそれをポインタ型として認識してくれる
// ただし、上記の話はポインタレシーバを持つメソッドの場合

// ポインタ変数
// a := "a"
// ptrA = &a
// &aはメモリアドレスを示しているもの
// *ptrA => "a"
// *をつけることで間接参照でき、変数の中身を見に行ける
