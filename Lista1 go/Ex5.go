package main

import "fmt"

var consumidor string
var Nconta int
var valor, agua float64

func main() {
	fmt.Println("Digite, separados por espaço, o número da conta, consumo de água em m³ e tipo de consumidor (R para residencial | C para comercial | I para industrial)")
	fmt.Scan(&Nconta, &agua, &consumidor)
	if consumidor == "r" || consumidor == "R" {
		valor = agua*0.05 + 5
		fmt.Println("Conta =", Nconta)
		fmt.Printf("Valor da conta= %.2f", valor)
	}
	if consumidor == "c" || consumidor == "C" {
		valor = (agua-80)*0.25 + 500
		fmt.Println("Conta =", Nconta)
		fmt.Printf("Valor conta= %.2f", valor)
	}
	if consumidor == "i" || consumidor == "I" {
		valor = (agua-100)*0.04 + 800
		fmt.Println("Conta =", Nconta)
		fmt.Printf("Valor da conta= %.2f", valor)
	}
}
