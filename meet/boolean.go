package meet

import "fmt"

func TypeBoolean() {
	var maior bool = 10 > 5
	var menor bool = 10 < 5
	fmt.Println("10 é maior que 5?", maior)
	fmt.Println("10 é menor que 5?", menor)
}

// Explicação detalhada
// var maior bool = 10 > 5
// Aqui você está comparando se 10 é maior que 5.
// O operador > retorna true porque a condição é verdadeira.
// Então a variável maior recebe o valor true.

// var menor bool = 10 < 5
// Agora você compara se 10 é menor que 5.
// O operador < retorna false, já que a condição não é verdadeira.
// A variável menor recebe o valor false.

// fmt.Println(...)
// Imprime as frases junto com os valores booleanos.
// A saída será:

// Código
// 10 é maior que 5? true
// 10 é menor que 5? false
// O que aprendemos aqui
// O tipo bool em Go só pode ser true ou false.

// Comparações (>, <, ==, !=, >=, <=) sempre retornam um valor booleano.

// Esses valores podem ser armazenados em variáveis, usados em condições (if, for) ou impressos diretamente.
