package main

import "fmt"

var n float64
var c string

func main() {
	fmt.Println("Digite a nota do aluno")
	fmt.Scan(&n)

	if n < 6 && n >= 0 {
		fmt.Printf("Nota = %.2f Conceito = D", n)
	} else if n >= 6 && n < 7.5 {
		fmt.Printf("Nota = %.2f Conceito = C", n)
	} else if n >= 7.5 && n < 9 {
		fmt.Printf("Nota = %.2f Conceito = B", n)
	} else if n >= 9 && n <= 10 {
		fmt.Printf("Nota = %.2f Conceito = A", n)
	} else {
		fmt.Println("Número inválido")
		return
	}
}
