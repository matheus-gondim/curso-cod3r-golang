package main

import "fmt"

func main() {
	var a int     // inicia com valor 0
	var b float64 // inicia com valor 0
	var c bool    // inicia com valor false
	var d string  // inicia com valor ""
	var e *int    // => ponteiro de inteiro // inicia com valor nil

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
