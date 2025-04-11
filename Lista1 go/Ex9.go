package main

import "fmt"

var a, b, c float64

func main() {
	fmt.Println("Digite os valores de A, B e C")
	fmt.Scan(&a, &b, &c)

	D := b*b - 4*a*c

	fmt.Printf("O valor de delta é = %.2f", D)
}
