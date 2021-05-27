package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor) se não tiver o * retorna o endereço na memoria
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
