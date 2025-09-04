package main

import "fmt"

func main() {
	r := Sum(20, 33)
	fmt.Println(r)
}

func Sum(n1 int, n2 int) int {
	r := n1 + n2
	return r
}
