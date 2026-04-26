package meet

import "fmt"

func TypeFloat() {
	// var floatNumber float32 = 1.1
	var pi float64 = 3.14       // constante matemática aproximada de π
	var raio float64 = 2.5      // raio do círculo
	var area = pi * raio * raio // fórmula da área: π * r²

	// fmt.Println("floatNumber:", floatNumber)
	fmt.Println("Raio:", raio)
	fmt.Println("Área do círculo:", area)
}

// Explicação detalhada
// var pi float64 = 3.14
// Aqui usamos float64, que é o tipo decimal de 64 bits. Ele oferece maior precisão (até ~15 casas decimais). É o tipo mais usado em cálculos matemáticos.

// var raio float64 = 2.5
// O raio também é float64, para manter consistência. Se você misturar float32 com float64, o Go exige conversão explícita.

// var area = pi * raio * raio
// Fórmula da área de um círculo: π × r².
// Como pi e raio são float64, o resultado também será float64.

// fmt.Println("Área do círculo:", area)
// Imprime o resultado no console. Nesse caso, 19.625.

// Sobre float32 vs float64
// float32 → ocupa menos memória, mas tem menos precisão (~7 casas decimais).

// float64 → ocupa mais memória, mas é muito mais preciso (~15 casas decimais).
// Por padrão, em cálculos matemáticos e científicos, usamos float64.

// Resumindo
// Esse exemplo mostra bem como usar float para cálculos com números decimais.

// float32 → útil em casos simples ou quando memória é crítica.

// float64 → recomendado na maioria dos casos, especialmente em cálculos matemáticos.
