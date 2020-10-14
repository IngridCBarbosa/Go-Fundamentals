package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")

	for i, aprovados := range aprovados {
		fmt.Printf("%d) %s\n", i, aprovados)
	}
}

func main() {
	aprovados := []string{
		"João Pedro",
		"Maria Luisa",
		"Joana",
		"Helena",
		"Jeronimo",
	}

	imprimirAprovados(aprovados...)
}
