package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2(parametro1 string, parametro2 string) {
	fmt.Printf("Funcao 2: %s %s\n", parametro1, parametro2)
}

func funcao3() string {
	return "Funcao 3"
}

func funcao4(parametro1, parametro2 string) string {
	return fmt.Sprintf("Funcao 4: %s %s", parametro1, parametro2)
}

func funcao5() (string, string) {
	return "retorno1", "retorno2"
}

func main() {
	funcao1()
	funcao2("parametro1", "parametro2")

	return3, return4 := funcao3(), funcao4("parametro1", "parametro2")
	fmt.Println(return3)
	fmt.Println(return4)

	return51, return52 := funcao5()
	fmt.Println(return51)
	fmt.Println(return52)
}
