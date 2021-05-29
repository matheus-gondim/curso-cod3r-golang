package main

import (
	"fmt"
	"reflect"
)

// slice não é um array, ele é um pedaço de um array.
// ou seja ele não cria um outro array, ele cria uma estrutura que tem um ponteiro
// que referencia a primeira posição e quantas posições ele deve ter
// se eu manipular o slice vai alterar no array principal
func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice => estrutura com tamanho variavel

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // vai do indice um ate o indice 3, mas nao incluindo o indice 3
	fmt.Println(a2, s2)

	a2[1] = 10
	fmt.Println(s2)

	s3 := a2[:2] // vai do inicio do array ate o indeci 2
	fmt.Println(s3)

	// vc pode imaginar um slice como: tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
