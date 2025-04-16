package main

import (
	"fmt"
)

func main(){
	var age, ctrl int
	var heig, weig float64
	AGE := []int{}
	HEIG := []float64{}
	WEIG := []float64{}

	for i := 1; ;i++{
		fmt.Println("Digite a idade, Peso e Altura:")
		fmt.Scan(&age, &weig, &heig)
	
		AGE = append(AGE, age)
		HEIG = append(HEIG, heig)
		WEIG = append(WEIG, weig)
		
		fmt.Println("Deseja continuar? Digite 1 para sim ou aperte qualquer outra tecla para sair!")
		fmt.Scan(&ctrl)
		if  ctrl != 1 {break}
	}

	fmt.Printf("- a quantidade de pessoas com idade superior a 50 anos é; %d\n", qtdPlusAge(AGE, 50))
	fmt.Printf("- a média das alturas das pessoas com idade entre 10 e 20 anos é; %.2f\n", medHigBtwAge(AGE, HEIG, 10, 20))
	fmt.Printf("- a porcentagem de pessoas com peso inferior a 40 quilos entre todas as pessoas analisadas é; %.2f%\n", percWeigLess(WEIG, 40))


}

func qtdPlusAge(l []int, ageMin int)(int){
	var qtd int

	for _, v := range l{
		if v > ageMin{
			qtd++
		}
	}

	return qtd
}

func medHigBtwAge(lAge []int, lHeig []float64, A1 int, A2 int)(float64){
	var qtd int
	var sum float64

	for k, v := range lAge{
		if v >= A1 && v < A2 || v <= A2 && v > A1{
			sum += lHeig[k]
			qtd++
		}
	}

	if qtd != 0{
		return sum/float64(qtd)
	}

	return -1
}

func percWeigLess(lweig []float64, maxWeig float64)(float64){
	var qtd float64
	
	for _, v := range lweig{
		if v < maxWeig{ qtd++ }
	}

	return (qtd/float64(len(lweig)))*100
}