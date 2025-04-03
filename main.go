package main
import "fmt"

func main(){

tamanhosapato := []int {47, 35, 8, 99, 22}
tamanhosapato = append(tamanhosapato, 33, 37, 44)
fmt.Println(tamanhosapato, len(tamanhosapato), cap(tamanhosapato))
}