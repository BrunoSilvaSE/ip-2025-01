package main

import(
	"fmt"
)

func main(){
	var salario float64

	fmt.Println("Digite o valor do salário")
	fmt.Scan(&salario)

	if salario <= 300{
		fmt.Printf("Salario com reajuste = %.2f\n", salario*1.5)
	}else{
		fmt.Printf("Salario com reajuste = %.2f\n", salario*1.3)
	}
}