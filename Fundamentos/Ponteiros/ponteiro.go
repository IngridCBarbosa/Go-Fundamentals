package main

import "fmt"

func main() {
	i := 1
	// GO NÃO TEM ARITMÉTICA DE PONTEIROS

	// CRIANDO UM PONTEIRO

	var p *int = nil // nil == null que é o nulo

	p = &i /* PEGANDO O ENDEREÇO DE MEMÓRIA DA VARIÁVEL I, e atribuindo ao
	ponteiro p. */

	*p++ // PEGANDO O VALOR DA VARIÁVEL
	fmt.Println("*p:", *p)
	fmt.Println("P:", p)
	fmt.Println("i:", i)
	fmt.Println("&i:", &i)
}
