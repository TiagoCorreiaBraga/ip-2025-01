package main

import "fmt"

var n int

func main() {
	fmt.Println("Digite um número positivo maior que 1")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}

	var soma float64 = 0.0
	for i := 1; i <= n; i++ {
		soma += 1.0 / float64(i)
	}
	fmt.Printf("%.6f\n", soma)
}
