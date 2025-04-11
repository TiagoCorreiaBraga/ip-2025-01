package main

import "fmt"

var conta, salario, consumo, custokw, contaD float64

func main() {
	fmt.Println("Valor salário mínimo")
	fmt.Scan(&salario)

	fmt.Println("Consumo de energia em kw")
	fmt.Scan(&consumo)

	custokw = salario / 70
	conta = custokw * consumo
	contaD = conta * 9 / 10

	fmt.Printf("Custo por kW: R$ %.2f\n", custokw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", conta)
	fmt.Printf("Custo com desconto: R$ %.2f\n", contaD)
}
