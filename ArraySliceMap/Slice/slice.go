package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //ARRAY => TEM TAMANHO FIXO
	s1 := []int{1, 2, 3}  // SLICE => TAMANHO VARIADO

	fmt.Println(a1, s1)

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice não é um array. O slice define o pedaço de um array.
	// pega o elemento de indice 1 e vai até o indice 3, não incluindo o indice 3 no slice, ou seja, o que vai conter no slice será o
	// elemento 1 e 2.
	s2 := a2[1:3]
	fmt.Println("Array:", a2)
	fmt.Println("Slice:", s2)

	s3 := a2[:2] // Um novo slice que aponta para o mesmo array
	fmt.Println("Array 2 e slice 3")
	fmt.Println(a2, s3)

	/* Podemos imaginar um slice como: tendo um tamanho e um ponteiro para um
	elemento de um array.
	*/

	s4 := s2[:1]
	fmt.Println("Slice 2:", s2)
	fmt.Println("Slice 4:", s4)
}
