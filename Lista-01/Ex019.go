package main

import(
	"fmt"
)

func main(){
	var n int
	var sum float32

	fmt.Println("Digite um número inteiro, positivo e maior que 1.")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("O número inserido é inválido!")
		return
	}

	for i := 1; i <= n; i++{
		sum = sum + 1/float32(i)
	}

	fmt.Printf("%.6f\n", sum)
}