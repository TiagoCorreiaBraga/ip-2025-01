package main

import "fmt"

var n1, n2, n3 float64

func main() {
	fmt.Println("Escreva as notas do aluno separadas somente por espaço")
	fmt.Scan(&n1, &n2, &n3)

	media := (n1 + n2 + n3) / 3

	fmt.Printf("média = %.2f\n", media)
	if media >= 6 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
