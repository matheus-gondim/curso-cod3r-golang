package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// metodo para alterar um valor dentro da struct
// deve-se usar um ponteiro, porque devemos alterar
// na memoria
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	if len(partes) >= 1 {
		p.sobrenome = partes[len(partes)-1]
	} else {
		p.sobrenome = ""
	}
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Matheus Pereira Gondim")
	fmt.Println(p1.getNomeCompleto())

}
