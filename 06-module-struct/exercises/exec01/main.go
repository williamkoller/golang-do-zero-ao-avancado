package main

import (
	"fmt"
	"time"
)

func main() {
	var items []ItemDaCompra

	items = append(items, ItemDaCompra{Nome: "arroz"})
	items = append(items, ItemDaCompra{Nome: "feijao"})
	compra := Compra{
		Mercado: "mercado x",
		Data:    time.Now(),
		Items:   items,
	}

	fmt.Println(compra)
}
