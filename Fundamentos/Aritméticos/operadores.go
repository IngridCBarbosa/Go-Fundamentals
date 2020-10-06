package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicaão =>", a*b)
	fmt.Println("Módulo =>", a%b)

	//  bitwise
	fmt.Println("AND =>", a&b) // 11 AND 10 = 10
	fmt.Println("OR =>", a|b)  // 11 OR 10 = 11
	// OU EXCLUSIVO
	fmt.Println("XOR =>", a^b) // 11 XOR 10 = 01

	c := 3.0
	d := 2.0

	// OUTRAS OPERAÕES

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
