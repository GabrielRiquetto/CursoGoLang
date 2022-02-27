package main

import (
	"fmt"
	"os"
)

func main() {
	name := introduction()
	getResponse := showMenu(name)
	doSomething(getResponse)
}

func introduction() string {
	var name string
	fmt.Println("Qual é o seu nome?")
	fmt.Scan(&name)
	fmt.Println("Seja bem vindo(a)")
	return name
}

func showMenu(name string) int {
	var response int
	fmt.Println("O que você deseja", name, "?")
	fmt.Println("1 - Inciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Scan(&response)
	return response
}

func doSomething(command int) {
	switch command {
	case 1:
		fmt.Println("Iniciando Monitoramento!")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo, até mais!")
		os.Exit(0)
	default:
		fmt.Println("Não entendi o que você quis dizer... Desculpe!")
		os.Exit(-1)
	}
}
