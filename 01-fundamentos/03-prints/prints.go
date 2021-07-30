package main

import "fmt"

func main() {
	// imprime no console
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// imprime no console e pula uma linha
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	// em vez de imprimir no console ela vai retornar uma string
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f.\n", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	/*
		%d => serve para inteiro
		%#x => serve para imprimir um hexdecimal
		%b => serve para imprimir o numero em binário
		%f => serve para float
		%t => serve para boolean
		%s => serve para string
		%v => serve para varios tipos diferentes
	*/

	fmt.Printf("%d %f %.1f %t %s\n", a, b, b, c, d)
	fmt.Printf("%v %v %v %v\n", a, b, c, d)
}
