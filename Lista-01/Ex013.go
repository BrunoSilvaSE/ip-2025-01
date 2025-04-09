package main

import(
	"fmt"
)

func main(){
	var nota float32
	
	tab := map[string][2]float32{
		"A" : {9.0, 10.1}, // o ultimo valor sempre deve ser 0.1 a mais
		"B" : {7.5, 9},
		"C" : {6, 7.5},
		"D" : {0, 6},
	}

	fmt.Println("Digite a nota a ser convertida")
	fmt.Scan(&nota)

	if nota >= tab["A"][1]{
		fmt.Println("O valor inserido é inválido, por favor, digite apenas valores inferiores à", tab["A"][1])
		return
	}

	for a, b := range tab{
		if b[0] <= nota && nota < b[1] {
			fmt.Printf("Nota = %.1f Conceito = %s", nota, a)
		}
	}
}