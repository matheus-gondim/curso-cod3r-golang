package main

import "fmt"

func main() {
	// mapas devem ser inicializados
	var aprovados map[int]string = make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12423478978] = "Pedro"
	aprovados[89586778978] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	fmt.Println(aprovados[12345678978])

	aprovados[00000000000] = "Matheus"
	fmt.Println(aprovados[00000000000])
	// se eu passar um valor para a mesma chave
	// o valor antigo e substituido
	aprovados[00000000000] = aprovados[00000000000] + " Gondim"
	fmt.Println(aprovados[00000000000])

}
