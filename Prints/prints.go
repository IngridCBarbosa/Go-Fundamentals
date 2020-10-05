package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516
	// fmt.print("O valor de x é "+x) NÃO FUNCIONA EM GO, pois x é do tipo float e ele
	// não permite que uma string se concatende com um tipo diferente

	xs := fmt.Sprint(x) // retorna uma string
	fmt.Println("O valor de x é " + xs)

	// OUTRA FORMA DE IMPRIMIR VALORES
	fmt.Println("O valor de x é", x)

	// UTILIZANO O PRINTF
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t  %s", a, b, b, c, d)
}
