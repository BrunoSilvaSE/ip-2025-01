package main

import(
	"fmt"
)

func main(){
	// Declaração de variáveis
	var value, venda float64
	tab := map[string][3]float64{
		"V1" : {0, 10, 70}, // "nome da compra" : {Valor mínimo, Valor máximo, Percentual de lucro em porcentagem}
		"V2" : {10, 30, 50},
		"V3" : {30, 50, 40},
		"MAX" : {50, 0, 30}, // MAX define um valor que não tem teto, com isso o valor maximo deve ser igual a 0
	}

	// entrada das informações
	fmt.Println("Digite o valor da compra:")
	fmt.Scan(&value)

	// cheacagem de erro
	if value <= 0 {
		fmt.Println("O valor inserido é invalido ou nulo")
		return
	}

	// processamento
	for n, list := range tab{
		if n == "MAX" && list[0] <= value{
			venda = value * (1+list[2]/100)
			break
		}else if list[0] <= value && value < list[1]{
			venda = value * (1+list[2]/100)
			break
		}
	}

	//saída
	fmt.Println(venda)

}