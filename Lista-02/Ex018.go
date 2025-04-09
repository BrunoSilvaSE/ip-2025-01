package main

import (
	"fmt"
)

func main(){
	var weekday int
	var class string
	var price, desconto, acrescimo float64

	fmt.Println("Digite o numero que faz referencia ao dia da semana:")
	fmt.Scan(&weekday)

	fmt.Println("Qual o preço normal do DVD?")
	fmt.Scan(&price)

	fmt.Println("Qual a class do DVD?\nDigite 'C' para comum ou 'L' para lançamento")
	fmt.Scan(&class)

	if weekday == 2 || weekday == 3 || weekday == 5{
		desconto = price*0.4
	}
	if class == "L" {
		acrescimo = price*0.15
	}

	FinalPrice := price + acrescimo - desconto
	fmt.Printf("O preço final é de R$%.2f\n", FinalPrice)
}