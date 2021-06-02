// estudar a implementação dos metodos map, reduce e filter
// como exemplo de uso de função como parametro
package main

import (
	"fmt"
)

func mutiplicacao(a, b int) int {
	return a * b
}

// devese passar a assinatuira da função
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	result := exec(mutiplicacao, 5, 3)
	fmt.Println(result)
}
