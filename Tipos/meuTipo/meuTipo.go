package main

import "fmt"

type Nota float64

func (nota Nota) entre(inicio, fim float64) bool {
	return (float64(nota) >= inicio && float64(nota) <= fim)
}

func notaParaConceito(nota Nota) string {
	if nota.entre(9.0, 10.0) {
		return "A"
	} else if nota.entre(7.0, 9.0) {
		return "B"
	} else if nota.entre(5.0, 7.99) {
		return "C"
	} else if nota.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
