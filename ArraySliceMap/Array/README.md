# Arrays
Array em go é muito similar com as outras linguagens de programação.

No programa, muitas vezes precisamos guardar uma coleção de dados do **mesmo tipo**, como a lista das notas dos alunos de uma determinada matéria. Esse tipo de armazenamento é feito em um programa utilizando array. 

Mas, o que é um array ?

Um array é uma sequencia de tamanho fixo, que é utilizado para armazenar elementos do mesmo tipo na memória. **O array de tamanho fixo não são muito populares como o slice em go**.

Em um array, é permitido armazenar zero ou mais elementos. Os elementos de um array são indexados utilizando **[0]**. A posição no array não é a partir do 1 e sim do 0.

*Exemplo:*

                        |1|2|3|
            index ->     0 1 2   

## Acessando array

    Os arrays podem ser criados de duas maneiras diferentes em Go.

1- Utilizando a palavra reservada **var** 
  
  *Exemplo:*

  ```Go
        var nome_array[tamanho]Tipo

        // ou

        var nome_array[tamanho]Tipo{ item1, item2, item3}
  ```

* Observações importantes

    * Em go, arrays são imutáveis, então você pode utilizar **array[index]** sintax no lado esquerdo

    ```Go
        var nome_array[tamanho] = elemento


    /* => keyword */  var array[3]string // => tipo do array
    ```

    * Podemos acessar o elemento do array usando o index ou utilizando um loop.
    * Em go, o tipo de array é dimensional
    * O tamanho do array é fixo e imutável
    * É permitido armazenar elementos duplicados em um array

2-Utilização a declaração curta

Em go, os arrays podem ser declarados usando a declaração curta. É mais flexível que a declaração mostrada acima.

*Exemplo:*

```Go
    nomeArray := [tamanho]Tipo{item1, item2, item3}
    
    nomeArray := [4]string{"Array", "Em", "Go", "!"}
```

## Observações importantes sobre um array

* Se um array não for inicializado explicitamente, então por padrão o valor desse array é zero;

* Podemos verificar o tamanho do array utilizando o **len()**

    *Exemplo:*
    ```Go
        len(array)
    ```

* Se uma ellipses "..." estiver no local do tamanho de um array, então o tamanho do array é determinado pelos elementos inicializado

    *Exemplos:*
    ```Go
        vogais := [...]string{"a","e","i","o","u"}
    ```

* Um array é o tipo de valor, não tipo de referência. O que isso significa ? Que quando um array é passando para uma nova variável, as mudanças feitas no novo array não afetará o array original

    *Exemplo:*
    ```Go
        numeros := [...]int{100, 200, 300}

        novoArrayNumeros := numeros
        
        novoArrayNumeros[0] = 500 // Essa mudança só afetará o array novoArrayNumeros
    ```

* Se o tipo do elementos do array é comparável, então o tipo do array também é comparável. Logo podemos comparar dois arrays diretamente utilizando o operador **==**.

    *Exemplo:*
    ```Go
        array1 := [3]int{7, 8, 9}
        array2 := [...]int{7, 8, 9}
        array3 := [3]int{2, 4, 7}

        fmt.Println(array1 == array2) // true
        fmt.Println(array1 == array3) // false
        fmt.Println(array2== array3) // false
    ```

## Fazendo cópia de  um array

Em go não há uma função que faça uma copia de um array, mas essa cópia pode ser feita ao criar uma nova variável e passar o array por valor ou por referência.

Qual a diferença ?
        
    Ao fazer uma cópia por valor, ao modificar o original por exemplo, o array cópia não será modificado, ou seja, as ações não terão reflexo uma na outra.

    Ao fazer uma cópia por referência, ao modificar por exemplo, a cópia , o original também será modificado, ou seja, as ações são refletidas

```go
    original := [] int {1, 2, 3, 4}

    // Cópia por valor 
    arrCopia := original

    // Cópia por referência
    arrCopia := &original

```