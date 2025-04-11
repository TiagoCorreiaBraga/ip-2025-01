package main

import (
	"fmt"
	"math"
)

var h, a float64

func main() {
	fmt.Println("Digite a altura e a aresta da base da pirâmide, ambos em M")
	fmt.Scan(&h, &a)

	ab := (3 * a * a * math.Sqrt(3)) / 2
	v := ab / 3 * h

	fmt.Printf("O volume da pirâmide é = %.2f metros cubicos", v)
}
