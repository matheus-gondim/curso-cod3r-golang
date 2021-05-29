package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota int = int(n)
	// em go por padrão assim que a expressão atende um caso
	// o processamento do switch e finalizado
	switch nota {
	case 10:
		fallthrough // o fallthrough serve para indicar que o switch deve executar o proximo case
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}

}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
