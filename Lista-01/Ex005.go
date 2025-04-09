//5 Conta de Água
package main

import (
	"fmt"
	"strings"
)

func main(){
	//var
	var cont_id int
	var consumo, vlconta float32
	var tipo string

	//input
	fmt.Println(strings.Repeat("=", 120))
	fmt.Printf("Tipos de consumidores\n C - COMERCIAL\n I -INDUSTRIAL\n R - RESIDENCIAL\n")
	fmt.Println(strings.Repeat("=", 120))
	fmt.Println("Digite a conta do cliente, o consumo por metros cúbicos e o tipo do consumido:")
	fmt.Scan(&cont_id, &consumo, &tipo)

	//processamento
	if strings.ToUpper(tipo) == "R" {
		vlconta = 5 + 0.05*consumo
	}
	if strings.ToUpper(tipo) == "C" {
		vlconta = 500 + (consumo-80)*0.25
		if consumo-80 <= 0 {
			vlconta = 500
		}
	}
	if strings.ToUpper(tipo) == "I" {
		vlconta = 800 + (consumo-100)*0.04
		if consumo-100 <= 0 {
			vlconta = 800
		}
	}

	//output
	fmt.Printf("Conta = %d\n", cont_id)
	fmt.Printf("Valor da conta = %.2f\n", vlconta)
}