package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// nesse caso e como se tivesse passando true como expressão
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

}
