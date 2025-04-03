package main
import "fmt"

func main(){

melhorespessoasdomundo := []string {"fabiano", "gisele", "heitor", "batman", "shadow"}
nomes1e2 := melhorespessoasdomundo[0:2]
fmt.Println(nomes1e2)
nomes3e4 := melhorespessoasdomundo[3:]
fmt.Println(nomes3e4)
nomedomeio := melhorespessoasdomundo[2]
fmt.Println(nomedomeio)
}