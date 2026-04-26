package meet

import "fmt"

func TypeInts() {
	// "int" é um número inteiro cujo tamanho depende da arquitetura da máquina:
	// em sistemas 64 bits, geralmente é int64; em 32 bits, int32.
	var idade int = 30

	// "int32" é um inteiro fixo de 32 bits.
	// Intervalo: -2.147.483.648 até 2.147.483.647.
	var contador int32 = 2

	// "int8" é um inteiro de 8 bits.
	// Intervalo: -128 até 127.
	var indice int8 = 1

	// "float32" é um número de ponto flutuante com precisão simples (32 bits).
	// Aproximadamente 7 casas decimais de precisão.
	var altura float32 = 1.75

	// "float64" é um número de ponto flutuante com precisão dupla (64 bits).
	// Aproximadamente 15 casas decimais de precisão.
	var peso float64 = 70.5

	// Impressões
	fmt.Println("Idade:", idade)
	fmt.Println("Contador:", contador)
	fmt.Println("Índice:", indice)
	fmt.Println("Altura:", altura)
	fmt.Println("Peso:", peso)

	// Exemplo de inferência de tipo:
	// O Go deduz automaticamente o tipo com base no valor.
	texto := "Olá, mundo!"
	numero := 42
	fmt.Println("Texto:", texto)
	fmt.Println("Número:", numero)
}

// Resumo dos tipos numéricos em Go
// int → inteiro genérico, depende da arquitetura (32 ou 64 bits).

// int8, int16, int32, int64 → inteiros com tamanho fixo, úteis para controle de memória ou interoperabilidade.

// uint8, uint16, uint32, uint64 → versões sem sinal (apenas valores positivos).

// float32 → números decimais com precisão simples.

// float64 → números decimais com precisão dupla (mais usado no dia a dia).
