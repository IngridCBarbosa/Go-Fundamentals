package main

import "fmt"

type nota float64

func (n nota) entre(nota1, nota2 float64) bool {
	return (float64(n) >= nota1 && float64(n) <= nota2)
}

func notaParaConteito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 9.0):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConteito(9.8))
	fmt.Println(notaParaConteito(6.9))
	fmt.Println(notaParaConteito(2.1))
}
