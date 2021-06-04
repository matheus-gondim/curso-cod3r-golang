package main

import (
	"fmt"
	"time"
)

// O channel é um ponto de sincronismo (e parada),
// sem ele o código é finalizado sem esperar a execução das goroutines

//goroutime causa um assincronismo
//já o channel e o que da o sincronismo, baseado no recebimento dos dados

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo, ele é sincrono, enquanto um dado nao chega
// ele fica esperando os dados chegarem

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- base * 2 // enviando dados para o canal

	time.Sleep(time.Second)
	c <- base * 3

	time.Sleep(time.Second * 3)
	c <- base * 4
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
