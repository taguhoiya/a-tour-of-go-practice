package main

import ("fmt")

func Pic(dx, dy int) [][]uint8 {
	// 長さdyのsliceに要素が
	// [1, 2, 3] dy = 3
	// 長さdxのsliceを割り当てる→長さdxのスライスにする
	yary := make([][]uint8, dy)
	fmt.Println(yary)
	// y_array := make([]uint8, dy)
	for v := range yary {
		yary[v] = make([]uint8, dx);
		for x:= 0; x< dx; x++ {
			yary[v][x] = 2
		}
	}
	return yary
}

func main() {
	Pic(3, 4)
}


// Pic 関数を実装してみましょう。 このプログラムを実行すると、生成した画像が下に表示されるはずです。
// この関数は、長さ dy のsliceに、各要素が8bitのunsigned int型で長さ dx のsliceを割り当てたものを返すように実装する必要があります。
// 画像は、整数値をグレースケール(実際はブルースケール)として解釈したものです。
// 生成する画像は、好きに選んでください。例えば、面白い関数に、 (x+y)/2 、 x*y 、 x^y などがあります。
