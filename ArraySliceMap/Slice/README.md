# Resumo slice

## O que é slice

Slice é mais poderoso, flexível e conveniente que um array, e é uma estrutura de dados
de peso leve. Slice é uma sequencia de elemenos do mesmo tipo, não é permitido adicionar 
dados de tipo diferente no mesmo. É como um array que apresenta um valor indexado e tamanho,
porém o tamanho do slice é dinamico, não é fixo que nem o do array.

```go
package main

func main() {
    // SINTAX

    s1 := []int
    s2 := []int{}
    s3 := []int{1, 2, 3, 4}
}
```
## Os três componentes do slice

 * __Ponteiro__
    
    O ponteiro é utilizado para apontar o primeiro elemento de um array que é acessível pelo slice.
    Aqui não é necessário que o ponteiro aponte para o primeiro elemento do array.

* __Length__
    
    Lenght é o numero total de elementos presento no array.

* __Capacity__
    
    Capacidade é o tamanho máximo que ele pode expandir. (É a memória que o array vai alocar).

## Como criar e inicializar o slice ?

 * __Usando Slice literal__
    
      Você pode criar um slice usando o slice literal. A criação do slice literal é como a do array literal,
      porém com uma diferença, não é permitido especificar o tamanho do slice nas chaves [].

      ```go
          var sliceLiteral = []string{"This","is", "slice","literal"}
      ```
 * __Usando um array__
    
      O slice é uma referencia do array, então podemos criar um slice de um dado array. Para criar o slice de um
      dado array primeiro temos que especificar [low:upper], que significa que o slice pode pegar os elementos começando 
      pelo lower até o upper. NÃO INCLUI O ELEMENTO UPPER.
    
    ```go
        array := [4]int{1, 2, 3, 4}

        slice1 := array[1:2]
        slice2 := array[0:]
        slice3 := array[:2]
        slice4 := array[:]

        fmt.Println("slice1:", slice1)
	    fmt.Println("slice2:", slice2)
	    fmt.Println("slice3:", slice3)
	    fmt.Println("slice4:", slice4)
    ```
    > slice1: [2]
    > slice2: [1,2,3,4]
    > slice3: [1,2]
    > slice4: [1,2,3,4]
 * __Utilizando um slice existente__
    
