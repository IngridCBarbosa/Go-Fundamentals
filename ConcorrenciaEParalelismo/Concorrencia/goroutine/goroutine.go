package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc não fala comigo", 3)
	// fale("João", "Só posso falar depois de vc", 1)

	//go fale("Maria", "Ei...", 500) // palavra reservada GO => significa que ele cria uma go routine, executa de forma independente
	// Como se fosse uma thread (lembrando que thread é uma linha de execução independente), porém mais leve.
	//go fale("João", "Opa...", 500)

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns", 5)
}
