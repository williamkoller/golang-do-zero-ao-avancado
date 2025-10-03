package model

import "time"

type Compra struct {
	Mercado string
	Data    time.Time
	Items   []ItemDaCompra
}

type ItemDaCompra struct {
	Nome string
}
