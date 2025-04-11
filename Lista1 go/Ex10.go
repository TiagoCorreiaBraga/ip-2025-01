package main

import "fmt"

var a, b, c, d float64

func main() {
	fmt.Println("Digite os valores de A, B, C e D")
	fmt.Scan(&a, &b, &c, &d)

	D := a*d - b*c

	fmt.Printf("O valor do determinante é = %.2f", D)
}
