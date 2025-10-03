package main

import (
	"fmt"
	"time"

	"github.com/williamkoller/06-module-struct-exercises01/model"
)

func main() {
	var items []model.ItemDaCompra

	items = append(items, model.ItemDaCompra{Nome: "arroz"})
	items = append(items, model.ItemDaCompra{Nome: "feijao"})
	compra := model.Compra{
		Mercado: "mercado x",
		Data:    time.Now(),
		Items:   items,
	}

	fmt.Println(compra)
}
