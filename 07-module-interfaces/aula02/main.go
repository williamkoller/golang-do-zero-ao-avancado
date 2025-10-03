package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error

	err = errors.New("teste")

	fmt.Println(err)

	ExibirError(errors.New("a error"))
	p := ProblemaDeNetwork{
		rede: true,
		hardware: false,
	}

	ExibirError(p)

}

func ExibirError(err error) {
	fmt.Println(err.Error())
}

type ProblemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	if p.rede {
		return "problema de rede"
	} else if p.hardware {
		return "problema de hardware"
	} else {
		return "sem erro"
	}

}
