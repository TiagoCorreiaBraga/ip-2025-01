package main

import "fmt"

var n int

func main() {
	fmt.Println("Digite o número")
	fmt.Scan(&n)

	if n <= 5 || n >= 2000 {
		fmt.Println("Valor fora do intervalo permitido (5 < N < 2000)")
		return
	}
	for i := 2; i <= n; i += 2 {
		fmt.Printf("%d^2 = %d\n", i, i*i)
	}
}
