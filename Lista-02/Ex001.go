package main

import(
	"fmt"
)

func main(){
	var n int

	fmt.Println("Digite um valor inteiro: ")
	fmt.Scan(&n)

	if n % 2 == 0{
		fmt.Println("O número digitado é par")
	}else{
		fmt.Println("O número digitado é ímpar")
	}
}