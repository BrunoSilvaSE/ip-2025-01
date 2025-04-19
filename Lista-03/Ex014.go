package main

import "fmt"

func main(){
	var n1, n2, divs int
	primos := []int{}

	fmt.Println("Qual deve ser o intervalo de busca?")
	fmt.Scan(&n1, &n2)

	if n1 > n2{
		constant := n1
		n1 = n2
		n2 = constant
	}

	for ;n1 <= n2; n1++{
		divs = 0
		for i := 1; i <= n1; i++{
			if n1 % i == 0{
				divs++
			}
		}
		if divs <= 2 && divs > 0{
			primos = append(primos, n1)
		}
	}

	fmt.Printf("Primos entre os valores inserido:\n")
	fmt.Println(primos)
}