package main

import "fmt"

var a1, r, n int

func main() {
	fmt.Println("Digite o termo inicial da PA, a razão da mesma e a quantidade de elementos")
	fmt.Scan(&a1, &r, &n)

	soma := 0
	for i := 0; i < n; i++ {
		termo := a1 + i*r
		soma += termo
	}
	fmt.Println(soma)
}
