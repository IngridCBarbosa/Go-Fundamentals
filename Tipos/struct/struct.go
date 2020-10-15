package main

import "fmt"

type Produto struct {
	// NÃO COLOCO VÍRGULA NA DEFINIÇÃO DA STRICT
	nome     string
	preco    float64
	desconto float64
}

// MÉTODO: função com receiver (receptor) => o receiver é quem recebe o método
// RECEPTOR(receiver)
func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 Produto

	// AQUI PRECISA DA VÍRGULA AO CRIAR UMA VARIÁVEL COM OS ATRIBUTOS
	produto1 = Produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	// FORMA REDUZIDA DE COMO CRIAR UMA VARIÁVEL DO TIPO STRUCT
	produto2 := Produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}
