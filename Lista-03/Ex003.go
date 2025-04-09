package main

import(
	"fmt"
)

func main(){
	var salarioC, salarioJ float64
	var tempo int
	taxC := 2.00
	taxJ := 5.00


	fmt.Println("\xF0\x9F\x92\xB3 Digite o salário de Carlos:")
	fmt.Scan(&salarioC)
	salarioJ = salarioC/3
	montC, montJ := salarioC, salarioJ

	for montC > montJ{
		montC *= 1+taxC/100
		montJ *= 1+taxJ/100
		tempo += 1
	}

	fmt.Printf("\xE2\x8C\x9B O tempo para que João passe Carlos no montante é de %d meses\n", tempo)
	fmt.Printf("Montante final\nJoão R$ %.2f\xF0\x9F\x92\xB0\n Carlos R$ %.2f\xF0\x9F\x92\xB0\n", montJ, montC)
}