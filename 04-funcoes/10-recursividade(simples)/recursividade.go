package main

import "fmt"

func fatorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {
	result := fatorial(5)
	fmt.Println(result)

}
