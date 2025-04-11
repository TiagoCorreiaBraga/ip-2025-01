package main

import "fmt"

var r, h float64

func main() {
	fmt.Println("Digite o raio e a altura da lata")
	fmt.Scan(&r, &h)

	π := 3.14159
	Ac := π * r * r
	Al := 2 * π * r * h
	At := 2*Ac + Al
	custo := At * 100

	fmt.Printf("o valor do custo é = %.2f", custo)
}
