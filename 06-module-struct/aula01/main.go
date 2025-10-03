package main

import "fmt"

type endereco struct {
	rua    string
	numero int
	cidade string
}

func main() {
	endereco := endereco{
		rua:    "Rua x",
		numero: 15,
		cidade: "Curitiba",
	}

	fmt.Println(endereco)
}
