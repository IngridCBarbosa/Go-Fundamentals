package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//  CUIDADO AO CONCATENAR STRINGS

	fmt.Println("Teste " + string(123)) // NÃO CONVERTE ESSE INTEIRO EM STRING

	// INT PARA STRING
	fmt.Println("Teste2 " + strconv.Itoa(123))

	// STRING PARA INT
	// O _ É QUE ESSE MÉTODO STRCONV RETORNA DOIS VALORES, O NUMERO E O ERRO. E O _
	// É QUANDO IGNORO O ERRO
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// STRING PARA BOOLEAN
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}

}
