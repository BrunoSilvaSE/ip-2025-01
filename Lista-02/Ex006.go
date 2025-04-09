package main

import(
	"fmt"
)

func main(){
	var nA, nB int

	fmt.Println("Digite em sequência o numerador e o denominador!")
	fmt.Scan(&nA, &nB)

	if nA % nB == 0{
		fmt.Printf("%d é divisível por %d\n", nA, nB)
	}else{
		fmt.Printf("%d não é divisível por %d\n", nA, nB)
	}
}