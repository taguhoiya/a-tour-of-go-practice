package main

import (
	"fmt"
)

type Vehicle interface {
	Accelerate()
	Brake()
}
type Car struct {
	color string
}

func (c Car) Accelerate() {
	fmt.Println("車が加速する")
}

func (c Car) Brake() {
	fmt.Println("ブレーキをかける")
}

type Bike struct {
	color string
}

func (c Bike) Accelerate() {
	fmt.Println("バイクが加速する")
}

func (b Bike) Brake() {
	fmt.Println("ブレーキをかける")
}

func drive(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Brake()
}

func main() {
	var car Car = Car{}
	drive(car)

	var bike Bike = Bike{}
	drive(bike)
}

// インターフェースとは
// 設計図（構造体）の共通の振る舞い（メソッド）を規約として定めたもの
// 定めたのち、インターフェースをレシーバの型として使うことで、構造体ごとに同じ振る舞いを書いたメソッドなどを必要がなくなる
