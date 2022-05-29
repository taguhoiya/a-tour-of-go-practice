package main

import (
	"fmt"
	"strings"
	// "sort"
)

func WordCount(s string) map[string]int {
	// ary := strings.Fields(s) //["I" "am" "learning" "Go!"]
	// sort.Strings(ary)
	exe := make(map[string]int)
	//[ "I" "ate" "a" "donut." "Then" "I" "ate" "another" "donut."]
	for _, v := range strings.Fields(s) {
		// exe[v]はゼロ値
		exe[v] = exe[v] + 1
	}
	return exe

}
func main() {
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
