package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Items   []ItemDaCompra
}

type ItemDaCompra struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, nomeDosItems []string) (*Compra, error) {

	if mercado == "" {
		return nil, errors.New("mercado é obrigatorio")
	}

	if len(nomeDosItems) == 0 {
		return nil, errors.New("items sao obrigatorios")
	}

	var items []ItemDaCompra

	for _, nome := range nomeDosItems {

		items = append(items, ItemDaCompra{Nome: nome})
	}

	return &Compra{
		Mercado: mercado,
		Data:    data,
		Items:   items,
	}, nil
}
