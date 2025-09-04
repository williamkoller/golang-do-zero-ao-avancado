package main

import "fmt"

func main() {
	soma, sub, div, mult := Operacao(100, 50)
	fmt.Println(soma)
	fmt.Println(sub)
	fmt.Println(div)
	fmt.Println(mult)
}

func Operacao(n1 int, n2 int) (int, int, int, int) {
	soma := n1 + n2
	sub := n1 - n2
	div := n1 / n2
	mult := n1 * n2
	return soma, sub, div, mult
}
