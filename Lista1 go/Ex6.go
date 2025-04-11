package main

import "fmt"

var tF [1000]float64
var tC [1000]float64
var medicoes int

func main() {
	fmt.Println("Digite a quantidade de medições")
	fmt.Scan(&medicoes)

	for i := 0; i < medicoes; i++ {
		fmt.Printf("Temperatura N.%d a ser convertida C\n", i+1)
		fmt.Scan(&tF[i])

		tC[i] = 5 * (tF[i] - 32) / 9
	}
	for i := 0; i < medicoes; i++ {
		fmt.Printf("%.2f Fahrenheit equivale à %.2f Celsius\n", tF[i], tC[i])
	}
}
