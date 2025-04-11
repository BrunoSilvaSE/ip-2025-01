package main

import "fmt"

func main(){
	var code int
	var valor, valorFinal, parcela float64

	fmt.Println("Qual o valor de etiqueta?")
	fmt.Scan(&valor)

	for true {

		fmt.Println("Qual a condição de pagamento? Digite o código!")
		fmt.Println("[1] - À vista, dinheiro ou cheque, 10% de desconto")
		fmt.Println("[2] - À vista, cartão de crédito, 5% de desconto")
		fmt.Println("[3] - em 2 vezes, preço normal de etiqueta sem juros")
		fmt.Println("[4] - em 3 vezes, preço normal de etiqueta + 10% de juros")
		fmt.Scan(&code)


		switch code {
		case 1:
			valorFinal = valor - valor*0.1
		case 2:
			valorFinal = valor - valor*0.05
		case 3:
			valorFinal = valor
			parcela = valorFinal/2
		case 4:
			valorFinal = valor + valor*0.1
			parcela = valorFinal/3
		default:
			fmt.Printf("O valor inserido '%d', é inválido\n", code)
		}

		if valorFinal != 0 {break}
	}

	fmt.Printf("Valor final a ser pago: R$ %.2f\n", valorFinal)
	if parcela != 0 {fmt.Printf("Valor por parcela: R$ %.2f\n", parcela)}
}