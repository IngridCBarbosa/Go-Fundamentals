package main

import "fmt"

func main() {
	/* QUANDO O ARRAY É DEFINIDO DESSA FORMA, QUEM CONTA A QUANTIDADES DE ELEMENTOS QUE O ARRAY TEM É O COMPILADOR, DAÍ DEFINE O TAMANHO DO ARRAY*/
	numeros := [...]int{1, 2, 3, 4, 5} // compilador que vai contar

	// O I: ÍNDICE, NUMERO: DADOS DO ARRAY
	/* SEMPRE RETORNA O INDICE E O NUMERO, SE POR ACASO QUISER IGNORAR O INDICE, É COLOCAR O _ NO LUGAR O I*/
	for i, numero := range numeros {
		fmt.Printf("%d) %d \n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
