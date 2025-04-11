package main

import "fmt"

func main() {

	var popular [1000]float64
	var geral [1000]float64
	var arquibancada [1000]float64
	var cadeiras [1000]float64
	var soma [1000]float64
	var pessoas [1000]float64
	var jogos int

	fmt.Println("Número de jogos")
	fmt.Scan(&jogos)

	for i := 0; i < jogos; i++ {
		fmt.Printf("Dados do jogo %d:\n", i+1)
		fmt.Scan(&pessoas[i])
		fmt.Scan(&popular[i])
		fmt.Scan(&geral[i])
		fmt.Scan(&arquibancada[i])
		fmt.Scan(&cadeiras[i])

		popular[i] = popular[i] / 100
		geral[i] = geral[i] / 100
		arquibancada[i] = arquibancada[i] / 100
		cadeiras[i] = cadeiras[i] / 100

		soma[i] = pessoas[i] * (popular[i] + geral[i]*5 + arquibancada[i]*10 + cadeiras[i]*20)
	}

	fmt.Println()

	for i := 0; i < jogos; i++ {
		fmt.Printf("A renda do jogo N. %d é = %.2f\n", i+1, soma[i])
	}
}
