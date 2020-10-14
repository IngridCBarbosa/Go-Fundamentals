package main

import "fmt"

func medias(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f\n", medias(7.7, 8.1, 5.9, 9.9))
}
