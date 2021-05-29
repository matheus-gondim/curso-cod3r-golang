package main

import (
	"fmt"
	"time"
)

func main() {

	// não tem while em Go, mas podemos usar o for desta forma para emular um while
	// while é usado para quando a gente tem uma loop com uma quantidade indeterminada
	// de repetição
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// for tradicional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	fmt.Println()

	for {
		// laço infinito
		fmt.Println("Para sempre...")
		time.Sleep(time.Second) // time.Second => representa 1 segundo
	}
}
