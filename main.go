package main

import (
	"fmt"
)

func main(){
    
    var numeros [5]int
    var soma int

    fmt.Println("Digite 5 numeros inteiros:")
    for i := 0; i < 5; i++ {
        fmt.Println("Numero %d: ", i+1)
       fmt.Scan(&numeros[i])
        soma += numeros[i]
    }

    fmt.Println("A soma dos números é: %d\n", soma)

}
