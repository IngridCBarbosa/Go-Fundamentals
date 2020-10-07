package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Não sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(time.Now()))
}
