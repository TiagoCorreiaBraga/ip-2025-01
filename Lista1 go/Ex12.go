package main

import "fmt"

var t int
var custo float64

func main() {
	fmt.Println("Tempo de uso da charrete em horas")
	fmt.Scan(&t)

	if t < 3 {
		custo = float64(t) * 5
		fmt.Printf("O valor a pagar é = %.2f", custo)
	} else {
		blocos := t / 3
		resto := t % 3

		custo = float64(blocos) * 10.0
		custo += float64(resto) * 5.0
		fmt.Printf("O valor a pagar é = %.2f", custo)
	}
}
