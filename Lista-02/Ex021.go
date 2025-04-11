package main

import (
	"fmt"
	"strings"
)

func main(){
	var n1, n2, n3, mEx float64
	var id int
	var conc, status string
	conceitos := map[string][2]float64{
		"A": {9, 10},
		"B": {7.5, 9},
		"C": {6, 7.5},
		"D": {4, 6},
		"E": {-1, 4},
	}

	fmt.Println("Qual o número de identificação do aluno?")
	fmt.Scan(&id)

	fmt.Println("Quais são as 3 notas do aluno?")
	fmt.Scan(&n1, &n2, &n3)

	fmt.Println("Qual a média dos exercícios do aluno?")
	fmt.Scan(&mEx)

	media := (n1+n2*2+n3*3+mEx)/7
	
	for k, v := range conceitos{
		if v[0] <= media && media >= v[1]{
			conc = string(k)
			break
		}
	}

	if conc == ""{
		fmt.Println("O conceito não foi encontrado")
		return
	}

	switch conc{
	case "A", "B", "C":
		status = "Aprovado \xE2\x9C\x85"
	case "D", "E":
		status = "Reprovado \xE2\x9D\x8C"
	default:
		status = ""
	}
	
	fmt.Println(strings.Repeat("*", 120))
	fmt.Printf("\xF0\x9F\x91\xA4 Aluno: %d\n", id)
	fmt.Printf("\xE2\x9C\x8F Média das notas: %.2f  %s\n", media, conc)
	fmt.Printf("\xF0\x9F\x93\x81 Status: %s\n", status)
	fmt.Printf("\xF0\x9F\x93\x91 Nota 1: %.2f | Nota 2: %.2f | Nota 3: %.2f\n", n1, n2, n3)
	fmt.Printf("\xF0\x9F\x93\x9A Média Exercícios: %.2f\n", mEx)
}