package main

import (
	"fmt"

)
func main(){
    estoque := map[string]int{
        "COXINHA": 10,
      "PÃO DE QUEIJO": 15,
      "REFRIGERANTE": 20,
        }
        fmt.Println("Estoque:")
        for produto, quantidade := range estoque {
          fmt.Printf("%s: %d\n", produto, quantidade)
        }
        estoque["COXINHA"] -= 2
        estoque["PÃO DE QUEIJO"] -= 1
      
        fmt.Println("Estoque atualizado:")
        for produto, quantidade := range estoque {
          fmt.Printf("%s: %d\n", produto, quantidade)
        }
      }