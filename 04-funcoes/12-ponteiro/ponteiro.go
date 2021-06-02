// EVITE DE PASSAR A REFERENCIA DE UMA VARIÁVEL
package main

import "fmt"

func inc1(n int) {
	n++
}

// um ponteiro é representado por um *
func inc2(n *int) {
	// o * na frente da variável é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // passando o valor da variável
	fmt.Println(n)

	// & a frente da variável é usado para obter o endereço da variável
	inc2(&n) // passando a referência
	fmt.Println(n)
}
