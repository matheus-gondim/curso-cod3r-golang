package main

import "fmt"

func main() {
	// [...] indica pro compilador que o array tera n posições
	// n sendo a quantidade de valores passados na inicialização
	numeros := [...]int{1, 2, 3, 4, 5}

	// foreach => assim podemos pegar o index e o item da interação sem precisar do contador
	for ix, numero := range numeros {
		fmt.Printf("%d) %d\n", ix, numero)
	}

	// podemos ignorar o index usando o "_"
	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}
}
