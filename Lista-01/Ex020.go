package main

import(
	"fmt"
)

func main(){
	var h, m, s, Tt int

	fmt.Println("Digite em sequencia um valor inteiro para as horas, minutos e segundos")
	fmt.Scan(&h, &m, &s)

	Tt = s + (m*60) + (h*3600)

	fmt.Printf("O tempo em segundos é = %d\n", Tt)

}
