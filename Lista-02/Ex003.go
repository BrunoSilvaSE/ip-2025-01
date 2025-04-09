package main

import (
	"fmt"
)

func main(){
	var n1, n2 int

	fmt.Println("Digite dois valores inteiros:")
	fmt.Scan(&n1, &n2)

	if n1+n2 > 20{
		fmt.Printf("O resultado é %d \n", n1+n2+8)
	}else{
		fmt.Printf("O resultado é %d \n", n1+n2-5)
	}
}