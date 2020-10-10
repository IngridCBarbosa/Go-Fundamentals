package main

import "fmt"

func main() {
	funcESalario := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 20000.00,
		"Pedro Junior":   1200.00,
	}

	funcESalario["Geni Maria"] = 10000.00

	fmt.Println(funcESalario)

	for _, salario := range funcESalario {
		fmt.Println(salario)
	}
}
