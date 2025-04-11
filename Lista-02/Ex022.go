package main

import(
	"fmt"
)

func main(){
	var FuncionarioID int 
	var QtdHorasExtras, descontoINSS, descontoIR float64
	salarioMin := 788.00
	valorHoraExtra := 10.00

	fmt.Println("Digite em sequência a matrícula do funcionário e a quantidade de horas-extras trabalhadas")
	fmt.Scan(&FuncionarioID, &QtdHorasExtras)

	salarioHoraExtra := QtdHorasExtras*valorHoraExtra
	salarioBruto := 3*salarioMin+salarioHoraExtra
	if salarioBruto > 1500{
		descontoINSS = salarioBruto*0.12
	}
	if salarioBruto > 2000{
		descontoIR = salarioBruto*0.20
	}
	salarioLiquido := salarioBruto - descontoINSS - descontoIR
	
	fmt.Printf("Funcionário: %d\n", FuncionarioID)
	fmt.Printf("Salário Bruto: %.2f\n", salarioBruto)
	fmt.Printf("Salário Liquido: %.2f\n", salarioLiquido)
}