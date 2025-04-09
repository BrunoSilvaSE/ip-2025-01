package main

import(
	"fmt"
)

func main(){
	var age int
	tab := map[string][2]int{
		"Recém-nascido" : {0,2},
		"Criança" : {3, 11},
		"Adolescente" : {12, 19},
		"Adulto" : {20, 55},
		"Idoso" : {56, 999},//Adicione o valor 999 para o ultimo valor
	}

	fmt.Println("Olá! Porfavor, informe a sua idade:")
	fmt.Scan(&age)

	for clas, AGE := range tab{
		if age >= AGE[0] && age <= AGE[1] {
			fmt.Printf("Olá, com %d anos você está classificado como %s! \xF0\x9F\x98\x84\n", age, clas)
		}
	}
}