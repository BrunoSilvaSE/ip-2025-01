package main

import (
	"fmt"
)

func main(){
	var n int
	
	for true{//input
		fmt.Println("Digite um número de 3 dígitos")
		fmt.Scan(&n)
		
		//verificação
		if n > 99 && n < 1000 {
			break
		}else{
			fmt.Println("O valor digitado é inválido, porfavor digite um numero de apenas 3 dígitos")
		}
	}
	//Processamento
	Cn := ((n-n%10)%100)/10

	//Output
	fmt.Printf("O caracter da casa de dezenas de %d é %d\n",n, Cn)
}