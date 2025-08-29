package main

import "fmt"

// array vs slices

func main() {
	// var list = [10]int{1, 2, 3, 4, 5, 6, 7, 8} // is one array
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8} // is one slices

	list = append(list, 1)

	fmt.Printf("%T", list)
}
