package main

import (
	"fmt"
	"reflect"
)

func somar(a int, b int) int {
	fmt.Println(reflect.TypeOf(a))
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}
