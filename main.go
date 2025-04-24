package main

import (
	"fmt"

)
func analisarNotas(nota1, nota2 float64) (float64, string){
	media := (nota1 + nota2) / 2

	if media >= 6 {
		return media, "Aprovado"
} else
     {

		return media, "Reprovado"
	}
}

func main(){
	media, status := analisarNotas(7.5, 7.5)
	fmt.Printf("Média: %.2f - Status: %s\n", media, status)
}