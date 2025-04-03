package main

import "fmt"
func main(){

	
    var nomes = []string{"amanda", "ana", "julia" , "carol" , "priscila"}
    fmt.Println(nomes)
     nomesOne := nomes [:2]
     fmt.Println(nomesOne)
     nomesTwo := nomes[3:]
     fmt.Println(nomesTwo)
     rangeOne := nomes[2]
     fmt.Println(rangeOne)
}