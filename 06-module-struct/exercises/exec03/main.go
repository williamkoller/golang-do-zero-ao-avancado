package main

import (
	"fmt"
	"time"

	"github.com/williamkoller/06-module-struct-exercises03/model"
)

func main() {

	var nomeDosItems []string

	nomeDosItems = append(nomeDosItems, "Arroz")
	nomeDosItems = append(nomeDosItems, "Feijao")
	nomeDosItems = append(nomeDosItems, "Banana")

	compra, err := model.NewCompra("Mercado x", time.Now(), nomeDosItems)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(compra)
}
