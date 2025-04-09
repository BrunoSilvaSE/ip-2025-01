package main

import(
	"fmt"
)

func main(){
	var np, n int

	fmt.Println("Digite dois valores inteiros, sendo o primeiro Par")
	fmt.Scan(&np, &n)

	if np%2!=0{
		fmt.Println("O PRIMEIRO NUMERO NÃO É PAR")
		return
	}

	for i := 1; i <= n;{
		if np%2==0 {
			fmt.Printf("%d ", np)
			np = np+2
			i++
		}
	}
}