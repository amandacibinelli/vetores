package main

import (
	"fmt"

)
func dadosPessoa(nome string, idade int) (int, string) {
	if idade >= 18 {
		return idade, "voce eh maior de idade"
	}
	return idade, "voce eh menor de idade"
}

func main() {
	idade, status := dadosPessoa("Amanda", 16)
	fmt.Printf("Idade: %d, Status: %s\n", idade, status)
}