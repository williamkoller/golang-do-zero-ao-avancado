package main

import "fmt"

func main() {
	var listAll = []int{2, 10, 9, 4, 5, 8, 1, 3}

	indexRight := listAll[:3]
	indexLeft := listAll[4:]
	lastItem := listAll[len(listAll)-1:]

	fmt.Println(indexRight)
	fmt.Println(indexLeft)
	fmt.Println(lastItem)
}
