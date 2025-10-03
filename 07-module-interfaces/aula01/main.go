package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

type retangulo struct {
	largura, altura float64
}

type circulo struct {
	radius float64
}

func ExibirGeometria(g geometria) {
	fmt.Println(g.area())
}

func main() {
	retangulo := retangulo{
		largura: 1,
		altura:  2,
	}

	circulo := circulo{
		radius: 3,
	}

	ExibirGeometria(retangulo)
	ExibirGeometria(circulo)
}
