package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome      string
	sobrenome string
}

func (pessoa Pessoa) getNomeCompleto() string {
	return pessoa.nome + " " + pessoa.sobrenome
}

func (pessoa *Pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	pessoa.nome = partes[0]
	pessoa.sobrenome = partes[1]
}

func main() {
	pessoa1 := Pessoa{"Pedro", "Silva"}
	fmt.Println(pessoa1.getNomeCompleto())

	pessoa1.setNomeCompleto("Maria Luísa")
	fmt.Println(pessoa1.getNomeCompleto())
}
