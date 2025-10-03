package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	sliceString := []string{"A", "B", "C", "D"}
	newInts := reverseGenerics(slice)
	newStrings := reverseGenerics(sliceString)
	fmt.Println(newInts)
	fmt.Println(newStrings)
}

// func reverse(slice []int) []int {
// 	newInts := make([]int, len(slice))

// 	newIntsLen := len(slice) - 1
// 	for i := 0; i < len(slice); i++ {
// 		newInts[newIntsLen] = slice[i]
// 		newIntsLen--
// 	}

// 	return newInts
// }

func reverseGenerics[T int | string](slice []T) []T {
	newInts := make([]T, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}

	return newInts
}
