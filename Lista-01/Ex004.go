//4 Consumo de energia
package main

import(
	"fmt"
)

func main(){
	//Variáveis
	var salario int
	var qtKw float32

	//inputs
	fmt.Println("Digite o valor me reais do salario mínimo")
	fmt.Scan(&salario)
	fmt.Println("Agora digite a quantidade de KW gasta na residência")
	fmt.Scan(&qtKw)

	//processamento
	cpKw := float32(salario)*0.007
	consumo := float32(qtKw)*cpKw
	consumoDesconto := consumo-consumo*0.1

	//output
	fmt.Printf("Custo por KW: R$%.2f\n", cpKw)
	fmt.Printf("Custo do consumo: R$%.2f\n", consumo)
	fmt.Printf("Custo com desconto: R$%.2f\n", consumoDesconto)
}