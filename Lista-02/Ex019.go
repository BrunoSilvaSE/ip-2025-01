package main

import (
	"fmt"
	"math"
)

func main(){
	//declaração
	var volume, area, raio, altura float64
	var obj int

	//Input
	fmt.Println("Qual formra geometrica deseja calcular?")
	fmt.Println("Digite o numero que faz referencia a ela!")
	fmt.Println("[0] - Cone Reto")
	fmt.Println("[1] - Cilindro")
	fmt.Println("[2] - Esfera")
	fmt.Scan(&obj)

	//Processamento
	//Cone Reto
	if obj == 0{
		volume = (math.Pi*math.Pow(raio, 2)*altura)/3
		area = math.Pi*raio*math.Sqrt(math.Pow(raio, 2)+math.Pow(raio, 2))
	}
	//Cilindro
	if obj == 1{
		volume = math.Pi*math.Pow()
	}
	//esfera
}