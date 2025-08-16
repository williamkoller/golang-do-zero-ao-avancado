package main

import "fmt"

func main() {
	list := []int{4, 9, 6, 6}
	list = append(list, 8)

	listString := []string{"golang", "c#", "java"}
	listString = append(listString, "ruby")

	fmt.Println("List: ", list)
	fmt.Println("Lista pos1: ", list[0])
	fmt.Println("Lista pos2: ", list[1])
	fmt.Println("Lista pos3: ", list[2])
	fmt.Println("Lista string: ", listString)


	listMake := make([]int, 1)
	listMake[0] = 5
	listMake = append(listMake, 2)
	listMake = append(listMake, 3)

	fmt.Printf("%T\n", listMake)
	somaTotal := 0

	for i := 0; i < len(listMake); i++ {
		somaTotal += listMake[i]
	}
	fmt.Println("Média: ", somaTotal/len(listMake))

}
