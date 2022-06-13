package main

func myFunc() {
	i := 0
Here:   //この行の最初の単語はコロンを最後に持ってくることでタグとなります。
	println(i)
	i++
	goto Here   //Hereにジャンプします。
}

func main() {
	myFunc()
}