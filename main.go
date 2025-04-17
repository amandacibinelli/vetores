package main

import (
	"fmt"

)

func dividir(dividendo int,divisor int) (int, string) { 

if divisor ==0{
return 0, "erro na divisão por zero"
} 
return dividendo/divisor, "Sem erro"

}
 func main() {
 resultado, erro:= dividir(10,2)
 if erro != "sem erro" {
    fmt.Println(erro)
 } else {
    fmt.Println("O resultado da divisão eh:" , resultado)
 }
 }