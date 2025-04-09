package main

import(
	"fmt"
)

func main(){
	var Pcar float64
	var answ string
	additional := map[string]float64{
		"Ar Condicionado": 1750.00,
		"Pintura Metálica": 800.00,
		"Vidro Elétrico": 1200.00,
		"Direção Hidráulica": 2000.00,
	}
	var CARadd = make([]string, len(additional))

	fmt.Printf("\xF0\x9F\x9A\x98 Qual é o valor do carro?\n")
	fmt.Scan(&Pcar)

	for add, price := range additional{
		for true{
			fmt.Printf("Deseja adicionar %s (R$ %2.f) ao seu carro?\nDigite S para sim e N para não\n", add, price)
			fmt.Scan(&answ)
			if answ=="S" || answ=="s" || answ=="N" || answ=="n" {
				break
			}else{
				fmt.Printf("\xE2\x9D\x8C O valor '%s' é invalido! \xE2\x9D\x8C\n", answ)
			}
		}
		if answ == "S" || answ == "s" {
			Pcar += price
			CARadd = append(CARadd, add)
			fmt.Printf("\xE2\x9C\x85 Valor adicionado com sucesso! \xE2\x9C\x85\nValor total do carro: %.2f\n", Pcar)
		}else{
			fmt.Printf("Valor total do carro: %.2f\n", Pcar)
		}
	}

	fmt.Printf("\xF0\x9F\x9A\x98 Valor final do carro: R$ %.2f\nADICIONAIS\n", Pcar)
	for _, add := range CARadd{
		if add == ""{continue}
		fmt.Printf("— %s (R$ %.2f)\n", add, additional[add])
	}
	fmt.Printf("Obrigado por comprar conosco! \xF0\x9F\x98\x83\n")
}