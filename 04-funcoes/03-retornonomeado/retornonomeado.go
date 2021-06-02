package main

import "fmt"

// como o retorno e nomeado, não precisa colocar as variaveis na frente do return
// sendo que as variavais ja foram atribuidas valores
func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	x, y := troca(1, 2)
	fmt.Println(x, y)
}
