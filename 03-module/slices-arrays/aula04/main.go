package main

import "fmt"

// sublistas
func main() {
	var list = []int{2, 10, 9, 4, 5, 8, 1, 3}
	var listMenorQueCinco = make([]int, 0)

	for i := 0; i < len(list); i++ {
		if list[i] < 5 {
			listMenorQueCinco = append(listMenorQueCinco, list[i])
			fmt.Println("Posição[", i, "]", listMenorQueCinco)
		}
	}

	fmt.Println(listMenorQueCinco)
}
