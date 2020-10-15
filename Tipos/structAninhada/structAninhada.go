package main

import "fmt"

type Item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type Pedido struct {
	userID int
	itens  []Item
}

func (p Pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := Pedido{
		userID: 1,
		itens: []Item{
			Item{1, 2, 12.10}, // TENTO NÃO UTILIZAR DESSA FORMA PARA NAO GERAR CONFUSÃO
			Item{produtoID: 2, quantidade: 1, preco: 23.49}, // SEMPRE ESCREVA DESSA FORMA!!!
			Item{11, 100, 3.13},
		},
	}
	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
}
