package meet

import (
	"fmt"
	"strings"
)

func TypeString() {
	var hello string = "Olá, mundo!"
	var question string = "Como vai?"

	//Concatenação
	var meet = hello + question
	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet))
}

// Explicação detalhada
// var hello string = "Olá, mundo!"
// Cria uma string com a saudação inicial.

// var question string = "Como vai?"
// Outra string com uma pergunta.

// var meet = hello + question
// Usa o operador + para concatenar (juntar) as duas strings.
// Resultado: "Olá, mundo!Como vai?".

// fmt.Println(meet)
// Imprime a string concatenada.

// fmt.Println(strings.ToUpper(meet))
// Aqui entra o pacote strings.

// ToUpper() transforma todos os caracteres em maiúsculos.

// Resultado: "OLÁ, MUNDO!COMO VAI?".

// Funções úteis do pacote strings
// Além de ToUpper, você pode usar várias funções:

// strings.ToLower("Texto") → deixa tudo minúsculo.

// strings.Contains("Olá, mundo!", "mundo") → verifica se contém a palavra "mundo".

// strings.Split("a,b,c", ",") → divide a string em partes: ["a", "b", "c"].

// strings.TrimSpace(" texto ") → remove espaços extras no início e fim.

// Resumindo
// Strings podem ser concatenadas com +.

// O pacote strings oferece funções poderosas para manipulação de texto.
