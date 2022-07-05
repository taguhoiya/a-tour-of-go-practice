package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	pt, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("スタックトレースの取得失敗")
		return
	}
	funcName := runtime.FuncForPC(pt).Name()

	logs := fmt.Sprintf("%v:%v, %v(): ", file, line, funcName)
	log.SetPrefix(logs)
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
}
