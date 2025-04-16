package main

import "fmt"

func main(){
	var n int
	fact := 1

	fmt.Scan(&n)

	if n < 0{
		fmt.Println("O valor inserido é inválido!")
		return
	}

	if n == 0 {
		fmt.Printf("%d! é igual á %d\n", n, fact)
		return
	}

	for i := 1; i <= n; i++{
		fact *= i
	}

	fmt.Printf("%d! é igual á %d\n", n, fact)
}