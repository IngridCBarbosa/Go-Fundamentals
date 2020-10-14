package main

import "fmt"

func main() {
	funcionarioPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  15456.79,
			"Gustavo Pereira": 8456.78,
		},
		"J": {
			"Jose Luis":   6234.56,
			"Joana Brito": 20000.00,
		},
		"P": {
			"Pedro Junior": 1234.59,
		},
	}

	fmt.Println(funcionarioPorLetra)
	fmt.Println(funcionarioPorLetra["P"])
	fmt.Println("Excluindo a letra P")
	delete(funcionarioPorLetra, "P")
	fmt.Println(funcionarioPorLetra)

	for letra, funcionario := range funcionarioPorLetra {
		fmt.Printf("Letra: %s funcionário:%s\n", letra, funcionario)
	}
}
