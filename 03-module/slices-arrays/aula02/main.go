package main

import "fmt"

// criando lista
func main() {
	list := []int{4, 9, 6, 7}
	fmt.Println("List: ", list)
	fmt.Println("List: ", list[0])
	fmt.Println("List: ", list[1])
	fmt.Println("List lenght: ", len(list))

	list = append(list, 8)

	fmt.Println("List lenght: ", len(list))

}
