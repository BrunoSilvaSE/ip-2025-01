package main

import(
	"fmt"
	"math"
)

func main(){
	var h, a float64

	fmt.Println("Digite em sequência a altura da pirâmide e o comprimento de uma das arestas do hexagono")
	fmt.Scan(&h, &a)

	volume := float64((3*math.Pow(a, 2)*math.Sqrt(3)/2)*h)/3
	
	fmt.Printf("O volume da piramide é = %.2f metros cúbicos\n",volume)
}