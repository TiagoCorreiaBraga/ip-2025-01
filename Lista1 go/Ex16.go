package main

import "fmt"

var s float64

func main() {
	fmt.Println("Salário do funcionário")
	fmt.Scan(&s)

	if s <= 300 {
		sa := s * 1.5
		fmt.Printf("Salário com reajuste = %.2f", sa)
	} else if s > 300 {
		sa := s * 1.3
		fmt.Printf("Salário com reajuste = %.2f", sa)
	}
}
