package main

import "fmt"

func main() {

	// MAPS DEVEM SER INICIALIZADOS

	// var aprovados map[int]string => ASSIM SÓ CRIA

	aprovados := make(map[int]string)

	aprovados[1234567897] = "Maria"
	aprovados[9765432125] = "Pedro"
	aprovados[5679123458] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// PARA ACESSAR O ELEMENTO DO MAP
	fmt.Println(aprovados[5679123458])

	// PARA DELELTAR UM ELEMENTO NO MAP

	delete(aprovados, 9765432125)
}
