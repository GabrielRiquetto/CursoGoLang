package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := introduction()
	for {
		getResponse := showMenu(name)
		switch getResponse {
		case 1:
			startingMonitoring()
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

func startingMonitoring() {
	fmt.Println("Monitoranto...")
	site := "https://random-status-code.herokuapp.com/"
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site", site, "está com problema!", response.StatusCode)
	}
}
