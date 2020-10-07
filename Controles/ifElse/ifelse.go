package main

import "fmt"

func imprimirResultado(nota float64) {

	// A CONDIÇÃO NÃO É ENVOLVIDA ENTRE PARENTESE, MAS DEVE SEMPRE TER A CHAVE
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
