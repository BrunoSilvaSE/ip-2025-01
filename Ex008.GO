package main

import(
	"fmt"
)

func main(){
	var r, h, valor float32

	fmt.Println("Digite em sequencia o valor do raio e da altura da lata em metros")
	fmt.Scan(&r, &h)

	valor = 628.318*r*(r+h)

	fmt.Printf("O Valor do custo é = R$ %.2f", valor)
}