package main

import "fmt"

func main() {
	// chan(channel) e um tipo pertencente a linguagem go
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
