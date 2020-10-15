package main

import "fmt"

type Carro struct {
	nome            string
	velocidadeAtual int
}

type Ferrari struct {
	Carro       // CAMPOS ANÓNIMO. Não é herança mas uma composição!!
	turboLigado bool
}

func main() {
	ferrari := Ferrari{}

	ferrari.nome = "F40" // pelo fato de ter criado um campo anônimo, conseguimos acessar diretamente o atributo nome
	ferrari.velocidadeAtual = 0
	ferrari.turboLigado = false

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", ferrari.nome, ferrari.turboLigado)
	fmt.Println(ferrari)
}
