package main

import "fmt"

// a função init e sempre executada antes da propria função main
// ela serve para fazer configurações iniciais do arquivo
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
