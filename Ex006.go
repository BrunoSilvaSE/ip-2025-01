package main

import (
	"fmt"
	"strings"
)

func main(){
	//declaração de variáveis
	var nLinha int
	var temp = make(map[int][]float32)

	//inputs
	fmt.Println(strings.Repeat("==", 60))
	fmt.Println("Digite a quantidade de temperaturas em Fahrenheit a serem convertidas para Celsius ")
	fmt.Scan(&nLinha)
	fmt.Println(strings.Repeat("==", 60))
	for i := 1; i <= nLinha; i++{
		var Fahrenheit, Celsius float32

		fmt.Printf("Digite o %d° valor a ser convertido.\n", i)
		fmt.Scan(&Fahrenheit)
		Celsius = 5*(Fahrenheit-32)/9
		temp[i] = []float32{Fahrenheit, Celsius} 
		fmt.Println(strings.Repeat("==", 60))
	}
	for _, n := range temp{
		fmt.Printf("%.2f Fahrenheit equivale a %.2f Celsius.\n", n[0], n[1])
	}
	fmt.Println(strings.Repeat("==", 60))
}