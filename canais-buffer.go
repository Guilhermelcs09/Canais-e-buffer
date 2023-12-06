package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Ola mundo"
	canal <- "Gui"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem2)
	fmt.Println(mensagem)
}
