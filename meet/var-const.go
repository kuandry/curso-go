package meet

import "fmt"

func Variaveis() {
	// Variáveis
	var nome string = "Douglas"
	var idade int = 30
	var altura float64 = 1.75

	// Constantes
	const curso string = "Go - Introdutório"

	// Impressões
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Curso:", curso)

	// Exemplo de inferência de tipo
	texto := "Olá, mundo!"
	fmt.Println(texto)
}
