package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo de raio foi inferido pelo compilaodr

	// FORMA REDUZIDA DE CRIA UMA VERIÁVEL
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

	var e, f bool = true, false

	fmt.Print(e, f)

	g, h, i := 2, false, "epa"

	fmt.Print(g, h, i)
}
