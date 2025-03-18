package main

import(
	"fmt"
)

func main(){
	var A, B, C, DELTA float32

	fmt.Println("Digite o Valor de A, B e C em sequencia para calcular o valor de Delta")
	fmt.Scan(&A, &B, &C)

	DELTA = B*B-4*A*C

	fmt.Printf("O Valor de Delta é = %.2f", DELTA)
}