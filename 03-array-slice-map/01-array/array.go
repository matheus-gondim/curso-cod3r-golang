package main

import "fmt"

func main() {
	// os arrays são uma estrutura homogênea e estatica,
	// ou seja eles sempre tem valores do mesmo tipo
	// tendo um tamanho fixo de posições

	// um array sempre vem inicializado com os valores default do tipo
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média %.2f\n", media)
}
