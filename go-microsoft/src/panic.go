package main

import "fmt"

func main() {
	val := 0
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &val)

	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("You entered negative number")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
		if val > 0 {
			break
		}
	}
}

// 		for {
// 			if val == 0 {
// 				fmt.Println("You entered:", val)
// 				fmt.Println("0 is neither negative nor positive")

// 				// for input again
// 				fmt.Print("Enter number: ")
// 				fmt.Scanf("%d", &val)
// 			} else if val < 0 {
// 				error := fmt.Sprintf("You entered negative number: %v", val)
// 				panic(error)
// 			} else {
// 				fmt.Println("You entered:", val)
// 				break
// 			}
// 		}
// }
