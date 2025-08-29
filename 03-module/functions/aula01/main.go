package main

import "fmt"

func main() {
	ImprimeMsg("Ola")
}

func ImprimeMsg(msg string) {
	msg += ", bom dia!"
	fmt.Println(msg)
}
