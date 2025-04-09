package main

import (
	"fmt"
	"strings"
)

func main(){
	// var declaration
	var class string
	var consum, FinalPrice float64
	var PriceTab = map[string][3]float64{
		"Residencial": {5, 0, 0.05},
		"Comercial": {500, 80, 0.25},
		"Industrial": {800, 100, 0.04},
	}

	//input de valores
	for true {
		fmt.Println("Qual o tipo do consumidor?")
		fmt.Println("-Residencial\n-Comercial\n-Industrial")
		fmt.Scan(&class)
		class = strings.TrimSpace(class)

		ERROR := true
		for k := range PriceTab{
			if strings.ToUpper(class) == strings.ToUpper(k){
				ERROR = false
				break
			}
		}
		if ERROR {
			fmt.Println("O valor inserido é inválido!")
			fmt.Println(strings.Repeat("=", 120))
			continue
		}else{break}
	}
	fmt.Println(strings.Repeat("=", 120))
	for true{
		fmt.Println("Qual o consumo em m³?")
		fmt.Scan(&consum)

		if consum <= 0 {
			fmt.Println("O valor inserido é inválido!")
			fmt.Println(strings.Repeat("=", 120))
		}else{break}
	}
	fmt.Println(strings.Repeat("=", 120))
	//Proces
	for k, v := range PriceTab {
		if strings.ToUpper(class) == strings.ToUpper(k) {
			if consum > v[1] {
				FinalPrice = v[0] + (consum-v[1])*v[2]
			}else{FinalPrice = v[0]}
		}
	}

	//Return
	fmt.Printf("Valor da conta: R$ %.2f\n", FinalPrice)

}