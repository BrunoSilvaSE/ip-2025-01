package main

import(
	"fmt"
)

func main(){
	var a1, r, n, sum int

	fmt.Println("Digite em sequencia o primeiro termo da sua sequencia, a razão e a quantidade de termos")
	fmt.Scan(&a1, &r, &n)

	sum = a1
	for i := 1; i <= n; i++ {
		sum = sum + r
		fmt.Println(i,sum)
	}

	fmt.Println(sum)
}