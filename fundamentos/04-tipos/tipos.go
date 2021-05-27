package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// tipos numericos inteiros
	fmt.Println(1, 2, 1000)

	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	/*
		tipos inteiros sem sinal (só números positivos)
		uint8 => byte
		uint16 => sort
		uint32 => int
		uint64 => long
	*/

	var b byte = 255 // alias para uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	/*
		tipos inteiros com sinal (numeros positivos e negativos)
		int8 => byte
		int16 => sort
		int32 => int
		int64 => long
	*/

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	// char no go e o int32 ou seja o tipo rune
	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	/*
		números reais
		float32 => float
		float64 => double
	*/
	var x float32 = 49.99 // ou pode se fazer asssim => var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Matheus"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1)) //len() retorna o tamanho da string

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Matheus`
	fmt.Println("O tamanho da string é", len(s2))

}
