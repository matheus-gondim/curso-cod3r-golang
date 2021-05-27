// Programas executáveis iniciam pelo pacote main
package main

/*
Os códigos em Go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports
*/
import "fmt"

// A porta de entrada de um programa Go é a função main
func main() {
	fmt.Print("Primeiro ")
	fmt.Print("Programa!")

	/*
		Sobre comentários...
		1) Priorize código legível e faça comentários que agrega valor!
		2) Evite comentários óbvios
		3) Durante o curso abuse dos comentários
	*/

	/*
		Comandos

		godoc -http=:PORT => serve para habilitar a documentação offline do go na porta escolhida
		go env => traz informações sobre o hambiente que o go ta instalado
		go vet <arquivo> => e uma forma simplificada de detectar construções suspeitas
		go build => builda o programa
		go get => instala um pacote (tipo npm install)

	*/
}
