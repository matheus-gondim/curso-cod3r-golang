package main

import "fmt"

func main() {
	s := make([]int, 10) // cria um slice de 10 posições
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) // cria um slice de 10 posições, mas o array interno tera 20 posições

	fmt.Println(s, len(s), cap(s)) // cap() => retorna a capacidade do array interno

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// quando atinge a capacidade maxima do array interno, essa capacidade aumenta
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
