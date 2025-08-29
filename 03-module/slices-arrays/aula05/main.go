package main

import "fmt"

// index

func main() {
	var list = []int{2, 10, 9, 4, 5, 8, 1, 3}

	segundaList := list[:3]
	terceiraList := list[6:]
	ultimoList := list[len(list)-1:]
	fmt.Print(list)
	fmt.Print("\n", segundaList)
	fmt.Print("\n", terceiraList)
	fmt.Print("\n", ultimoList)
}
