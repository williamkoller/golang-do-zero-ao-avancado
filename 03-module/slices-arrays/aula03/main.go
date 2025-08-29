package main

import "fmt"

func main() {
	var listInt = []int{2, 10, 9, 4, 5, 8, 1, 3}
	var listMenorCinco = make([]int, 0)
	for i := 0; i < len(listInt); i++ {
		if listInt[i] < 5 {
			listMenorCinco = append(listMenorCinco, listInt[i])
			fmt.Println("posicao[", i, "]", listMenorCinco)

		}
	}

	fmt.Println(listMenorCinco)

}
