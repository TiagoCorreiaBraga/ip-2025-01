package main

import "fmt"

var n1, n2 int

func main() {
	fmt.Println("Digite os dois números pares")
	fmt.Scan(&n1, &n2)

	if n1%2 != 0 {
		fmt.Println("o primeiro número não é par")
		return
	}
	for i := 0; i < n2; i++ {
		fmt.Printf("%d ", n1+(i*2))
	}
}
