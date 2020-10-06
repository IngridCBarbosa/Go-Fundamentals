package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo de raio foi inferido pelo compilaodr

	// FORMA REDUZIDA DE CRIA UMA VERIÁVEL
	// QUANDO DEFINIDO DESSA FORMA É UMA VARIÁVEL E NÃO UMA CONSTANTE
	// 1- SEMPRE QUE DEFINIR UMA VARIÁVEL UTILIZA-LA
	area := PI * math.Pow(raio, 2)

	fmt.Print("o resultado da area é :", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// e = true e f = false
	var e, f bool = true, false

	fmt.Print(e, f)

	// g = 2 , h = false e i = "epa"
	g, h, i := 2, false, "epa"

	fmt.Print(g, h, i)
}
