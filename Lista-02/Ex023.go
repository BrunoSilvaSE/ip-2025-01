package main

import "fmt"

func main(){
	var age int

	fmt.Println("Escreva a sua idade: ")
	fmt.Scan(&age)

	if age < 16 && age >= 0{
		fmt.Println("Não-eleitor")
	}else if age >= 18 && age < 65{
		fmt.Println("Eleitor obrigatório")
	}else if age >= 16 && age < 18 || age >= 65{
		fmt.Println("Eleitor facultativo")
	}else{
		fmt.Println("O Valor digitado é inválido!")
	}

}

