package main

import (
	"fmt"
	"time"

	"github.com/williamkoller/06-module-struct-aula04/model"
)

func main() {
	endereco := model.Endereco{
		Rua:    "Rua x",
		Numero: 15,
		Cidade: "Curitiba",
	}
	pessoa := model.Pessoa{
		Nome:             "William",
		Endereco:         endereco,
		DataDeNascimento: time.Date(1989, 06, 21, 0, 0, 0, 0, time.Local),
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)
	endereco.Numero = 18
	fmt.Println(endereco.Numero)
	idade := model.CalculaIdade(pessoa)
	fmt.Println(idade)
	fmt.Println(pessoa.IdadeAtual())
}
