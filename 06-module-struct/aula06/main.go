package main

import (
	"fmt"
	"time"

	"github.com/williamkoller/06-module-struct-aula06/model"
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
	fmt.Println(endereco.Numero)
	pessoa.CalculaIdade()
	fmt.Println(pessoa.Idade)

	automodelMoto := model.Automovel{
		Ano:    2025,
		Placa:  "xpto",
		Modelo: "CG",
	}

	moto := model.Moto{
		Automovel:   automodelMoto,
		Cilindradas: 125,
	}

	fmt.Println(moto)

	automovelCarro := model.Automovel{
		Ano:    2025,
		Placa:  "xpto",
		Modelo: "BMW",
	}

	carro := model.Carro{
		Automovel:            automovelCarro,
		QuantidadeDePortas:   4,
		PossuiArCondicionado: true,
		Potencia:             2.0,
	}

	fmt.Println(carro)

}
