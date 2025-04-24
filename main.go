package main

import (
	"fmt"

)
func pegarNome() (string, string) {
    return "Barry" , "Allen"
}


func main(){
    nome, sobrenome := pegarNome()
        fmt.Printf(nome)
        fmt.Printf(sobrenome)

}
