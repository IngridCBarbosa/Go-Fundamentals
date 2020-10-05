package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// NUMÉRICOS INTEIROS
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// SÓ VALORES POSITIVOS  uint8 uint16 uint32 uint64 u(unsigned)

	var b byte = /* uint8 */ 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// COM SINAL int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // represneta um mapeamento da pabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// NUMÉROS REAIS
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// BOOLEAN
	bo := true
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// STRING

	s1 := "Olá eu sou uma string"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// STRING COM MUITAS LINHAS
	s2 := `Olá
	eu
	sou 
	uma
	string
	`
	fmt.Println("O tamanho da string é", len(s2))

	// CHAR => não exite o tipo char no GO, é int32
	// var x char = 'b' => não existe!

	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}
