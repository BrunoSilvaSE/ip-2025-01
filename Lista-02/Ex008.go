package main

import(
	"fmt"
)

func main(){
	var n int

	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&n)

	if n > 20 && n < 90{
		fmt.Printf("O número %d está entre 20 e 90\n", n)
	}else{
		fmt.Printf("O número %d não está entre 20 e 90\n", n)
	}
}