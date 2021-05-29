package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	fmt.Println("Antes de deletar um campo do map")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	delete(funcsPorLetra, "P")

	fmt.Printf("\nDepois de deletar um campo do map\n")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	fmt.Printf("\nPegando cada funcionário do map\n")
	for letra, funcs := range funcsPorLetra {
		fmt.Println("Funcionários com a letra", letra)

		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
