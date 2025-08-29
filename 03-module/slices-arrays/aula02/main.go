package main

import "fmt"

func main() {
	list := []int{4, 9, 6, 7}
	list = append(list, 8)
	liststring := []string{"golang", "java", "c#"}
	liststring = append(liststring, "python")
	fmt.Println("List: ", list)
	fmt.Println("List: ", list[0])
	fmt.Println("List: ", list[1])
	fmt.Println("List lenght: ", len(list))

	fmt.Println("List lenght: ", len(list))
	fmt.Println("list string: ", liststring)

	listMake := make([]int, 1)
	listMake[0] = 5
	listMake = append(listMake, 2)
	listMake = append(listMake, 3)
	fmt.Printf("%T\n", listMake)

	sumTotal := 0
	for i := 0; i < len(listMake); i++ {
		sumTotal += listMake[i]
	}
	fmt.Println("Media: ", sumTotal/len(listMake))
}
