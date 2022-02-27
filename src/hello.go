package main

import (
	"fmt"
	"os"
)

func main() {
	var command int

	introduction()
	showMenu()
	command = getCommand()
	doSomething(command)

}

func introduction() {
	var name string
	fmt.Println("Qual é o seu nome?")
	fmt.Scan(&name)

	fmt.Println("Olá, Sr.", name)

}

func showMenu() {
	fmt.Println("O que você deseja?")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func getCommand() int {
	var comando int

	fmt.Scan(&comando)
	return comando
}

func doSomething(command int) {
	if command == 1 {
		fmt.Println("Iniciando Monitoramento")
	} else if command == 2 {
		fmt.Println("Exibindo Logs")
	} else if command == 0 {
		fmt.Println("Saindo... Até mais!!")
		os.Exit(0)
	} else {
		fmt.Println("Não conheço o seu comando, desculpe =(")
		os.Exit(-1)
	}
}
