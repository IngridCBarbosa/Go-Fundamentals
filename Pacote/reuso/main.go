package main

import (
	"fmt"

	"github.com/IngridCBarbosa/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	// fmt.Println(area._TrianguloEq(5.0, 2.0)) => não funcionar , pois essa função é privada para outros pacotes
}
