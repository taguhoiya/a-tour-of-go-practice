package main

import "fmt"

func main() {
	var tax_rate float64
	var goods_a, goods_b, goods_c string
	var goods map[string]int = map[string]int{}
	//　{key:value} => {nil}

	tax_rate = 1.1
	goods_a = "大根"
	goods_b = "にんじん"
	goods_c = "トマト"
	goods[goods_a] = 200
	goods[goods_b] = 80
	goods[goods_c] = 150
	cost := calc(tax_rate, goods)
	fmt.Printf("総額%v円", cost)
}

func calc(tax_rate float64, goods map[string]int) float64 {
	total_cost := 0.0
	for _, v := range goods {
		total_cost += float64(v) * tax_rate
	}
	return total_cost
}
