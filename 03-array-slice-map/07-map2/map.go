package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	// excluir um documento que não existe no map nãp gera erro
	delete(funcsESalarios, "Inexistente")

	// em um for de map o primeiro valor sempre sera a chave
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
