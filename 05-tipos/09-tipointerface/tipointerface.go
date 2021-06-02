package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	// usando a interface vazia como um tipo, a variável se comporta como o any do typescript
	// recebe qualquer valor e ou tipo
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
}
