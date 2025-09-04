package main

import "fmt"

var n = 5

func main() {
	soma, sub, div, mult := Operacao(100, 50)
	soma -= n
	fmt.Println(soma)
	fmt.Println(sub)
	fmt.Println(div)
	fmt.Println(mult)
}

func Operacao(n1 int, n2 int) (soma int, sub int, div int, mult int) {
	soma = n1 + n2 + n
	sub = n1 - n2
	div = n1 / n2
	mult = n1 * n2
	return
}
