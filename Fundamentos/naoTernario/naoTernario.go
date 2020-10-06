package main

import "fmt"

// NÃO TEM O OPERADOR TERNÁRIO

func obterResultado(nota float64) string {

	if nota >= 7 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
