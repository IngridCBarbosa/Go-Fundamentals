package main

import "fmt"

func inc1(n int) {
	n++
}

// UM PONTEIRO É REPRESENTADO POR UM *

//  FUNCAO ABAIXO NÃO É UMA FUNÇÃO PURA
func inc2(n *int) {
	// * é usado para desrefenciar , ou seja, ter
	// acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // cópia do valor
	fmt.Println(n)

	// & USADO PARA OBTER O ENDEREÇO DA VARIÁVEL
	inc2(&n)
	fmt.Println(n) // PASSANDO O VALOR POR REFERÊNCIA
}
