package main

import "fmt"

var n1, n2, n3 int

func main() {
	fmt.Println("Digite os 3 números (de 1 dígito) inteiros separados por espaço")
	fmt.Scan(&n1, &n2, &n3)

	numero := n1*100 + n2*10 + n3
	quadrado := numero * numero

	if n1 > 9 || n2 > 9 || n3 > 9 {
		fmt.Println("Número inválido")
		return
	}
	fmt.Println(numero, ",", quadrado)
}
