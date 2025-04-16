package main

import "fmt"

func main(){
	var n, t1, t2, soma int
	t2 = 1

	fmt.Scan(&n)

	if n == 0{
		return
	}
	if n >= 1{
		fmt.Printf("%d-", t1)
	}
	if n >= 2{
		fmt.Printf("%d-", t2)
	}
	if n > 2{
		for i := 3; i <= n; i++{
			soma = t1 + t2
			t1 = t2
			t2 = soma
			fmt.Printf("%d-", soma)
		}
	}
	fmt.Println("")
}