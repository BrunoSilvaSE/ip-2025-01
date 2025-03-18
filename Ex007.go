package main

import(
	"fmt"
)

func main(){
	//var
	var Fahrenheit, Polegadas, Celsius, Milimetros float32

	//inputs
	fmt.Println("Digite um valor em Fahrenheit e outro valor em polegadas para que sejam devidamente convertidos!")
	fmt.Scan(&Fahrenheit, &Polegadas)

	//Processamento
	Celsius = 5*(Fahrenheit-32)/9
	Milimetros = Polegadas*25.4

	//output
	fmt.Printf("\nO VALOR EM CELSIUS = %.2f°C\n", Celsius)
	fmt.Printf("\nA QUANTIDADE DE CHUVA É = %.2fmm\n", Milimetros)

}