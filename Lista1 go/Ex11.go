package main

import "fmt"

var n int

func main() {
	fmt.Println("Digite o número")
	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 00 {
		fmt.Println("O número é divisível")
	} else {
		fmt.Println("O número não é divisível")
	}
}
