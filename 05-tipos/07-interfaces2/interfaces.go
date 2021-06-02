package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	tuboLigado      bool
	valocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.tuboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)

}
