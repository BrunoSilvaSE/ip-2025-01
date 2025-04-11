package main

import (
	"fmt"
	"math"
)

func main(){
	//declaração
	var volume, area, raio, altura float64
	var obj int
	objetos := [3]string{"Cone Reto", "Cilindro", "Esfera"}

	//Input
	fmt.Println("Qual formra geometrica deseja calcular?")
	fmt.Println("Digite o numero que faz referencia a ela!")
	fmt.Println("[0] - Cone Reto")
	fmt.Println("[1] - Cilindro")
	fmt.Println("[2] - Esfera")
	fmt.Scan(&obj)
	fmt.Println("Agora digite o valor do raio e altura em centímetros")
	fmt.Scan(&raio, &altura)


	//Processamento
	//Cone Reto
	if obj == 0{
		volume = (math.Pi*math.Pow(raio, 2)*altura)/3
		area = math.Pi*raio*math.Sqrt(math.Pow(raio, 2)+math.Pow(raio, 2))
	
	//Cilindro
	}else if obj == 1{
		volume = (math.Pi*math.Pow(raio, 2)*altura)/3
		area = 2*math.Pi*raio*altura
	
	//esfera
	}else if obj == 2{
		volume = (4/3)*math.Pi*math.Pow(raio, 3)
		area = 4*math.Pi*math.Pow(raio, 2)
	}

	//Output
	fmt.Printf("Objeto: %s\nVolume: %.2fcm³\nÁrea: %.2fcm²\n", objetos[obj], volume, area)
}