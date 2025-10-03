package main

import (
	"fmt"

	"github.com/williamkoller/06-module-struct-aula03/model"
)

func main() {
	endereco := model.Endereco{
		Rua:    "Rua x",
		Numero: 15,
		Cidade: "Curitiba",
	}
	pessoa := model.Pessoa{
		Nome:     "William",
		Endereco: endereco,
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)
	endereco.Numero = 18
	fmt.Println(endereco.Numero)
}
