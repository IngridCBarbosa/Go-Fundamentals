# Concorrência vs Paralelismo

## Paralelismo 
> * Significa fazer duas coisas ao mesmo tempo. 
> * No mundo dos computadores, ter processamento simultâneo. (múltiplos processadores)
> * Executar processos simultaneamente em processadores físicos diferentes.

 ## Concorrência 
   > * Administrar múltiplas tarefas, podendo ocorrer apenas de um único processador.
    
 > * Capacidade de intercalar **administrar** vários processos ao mesmo tempo e isso pode ocorrer em um único processador físico.

> * A concorrência viabiliza o paralelismo. É uma forma de estruturar (**modelar**) o seu programa.

  #### Observações:
  * É possível que a concorrência seja melhor que o paralelismo;  
  * O paralelismo é mais custoso que a concorrência, pois nem sempre adicionando mais processadores é a forma mais eficiente de ser feita;
   * Paralelismo exige muito mais do SO e do hardware. 

### Go é uma linguagem concorrente !