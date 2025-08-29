package main

import "fmt"

func main() {
	slice := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	firstNumns := 0
	secondNums := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] <= 5 {
			firstNumns += slice[i]
		} else {
			secondNums += slice[i]
		}
	}

	fmt.Println(firstNumns)
	fmt.Println(secondNums)
}
