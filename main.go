package main

import (
	"fmt"

)
func dadosPessoa(nome string, idade int) (int, string) {
	var status string
	if idade >= 18 {
		status = "Maior de idade"
	} else {
		status = "Menor de idade"
	}
	return idade, status
}

func main() {
	var nome string
	var idade int

	fmt.Print("Digite o nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite a idade: ")
	fmt.Scanln(&idade)

	idadeRetornada, status := dadosPessoa(nome, idade)

	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d\n", idadeRetornada)
	fmt.Printf("Status: %s\n", status)
}