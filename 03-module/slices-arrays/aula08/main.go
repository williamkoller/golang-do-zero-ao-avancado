package main

import "fmt"

func main() {
	// m := map[string]int{"sp": 10000000, "cg": 900000}
	mm := make(map[string]int)

	mm["cwb"] = 1000000
	mm["lon"] = 10000
	for i, v := range mm {
		fmt.Println("Cidade", i, "habitantes", v)
	}
}
