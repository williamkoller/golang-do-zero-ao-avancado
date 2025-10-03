package main

import (
	"fmt"

	"github.com/williamkoller/06-module-struct-aula02/model"
)

func main() {
	endereco := model.Endereco{
		Rua:    "Rua x",
		Numero: 15,
		Cidade: "Curitiba",
	}

	fmt.Println(endereco)
	endereco.Numero = 18
	fmt.Println(endereco.Numero)
}
