package main

import(
	"fmt"
)

func main(){
	var n int

	fmt.Println("Digite um valor inteiro! ")
	fmt.Scan(&n)

	if n > 0{
		fmt.Println("O valor digitado é positivo")
	}else if n < 0{
		fmt.Println("O valor digitado é negativo")
	}else{
		fmt.Println("O valor digitado é nulo")
	}
}