package main

import "fmt"

func main() {
	var list []interface{}

	list = append(list, 10)
	list = append(list, true)
	list = append(list, "teste")

	for _, valor := range list {
		if v, ok := valor.(string); ok {
			fmt.Println(v + " string")
		} else {
			fmt.Println(valor)
		}
	}
}
