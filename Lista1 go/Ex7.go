package main

import "fmt"

var tF, qP float64

func main() {
	fmt.Println("Digite a temperatura em F a ser convertida em C e a quantidade de polegadas a ser convertida em mm")
	fmt.Scan(&tF, &qP)

	tC := (5*tF - 160) / 9
	qmm := qP * 25.4

	fmt.Printf("%.2f F em Celsius = %.2f\n", tF, tC)
	fmt.Printf("%.2f Polegadas em mm = %.2f", qP, qmm)
}
