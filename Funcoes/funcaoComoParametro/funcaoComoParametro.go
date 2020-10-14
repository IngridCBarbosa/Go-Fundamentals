package main

import "fmt"

func multipicao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, parametro1, parametro2 int) int {
	return funcao(parametro1, parametro2)
}

func main() {
	resultado := exec(multipicao, 4, 3)
	fmt.Println("Resultado:", resultado)
}
