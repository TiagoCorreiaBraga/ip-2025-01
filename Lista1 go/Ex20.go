package main

import "fmt"

var h, m, s int

func main() {
	fmt.Println("Digite o horário separado apenas por espaços")
	fmt.Scan(&h, &m, &s)

	ts := float64(s) + float64(m)*60 + float64(h)*60*60

	fmt.Printf("O tempo em segundos é = %.2f", ts)
}
