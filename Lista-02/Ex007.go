package main

import(
	"fmt"
)

func main(){
	var A, B, C float64
	var MAIOR, MENOR, INTER float64

	fmt.Println("Digite 3 números:")
	fmt.Scan(&A, &B, &C)

	// Definindo a ordem
	if A >= B && A >= C{
		MAIOR = A
		if C >= B{
			INTER, MENOR = C, B
		}else{
			INTER, MENOR = B, C
		}
	}else if B >= C && B >= C{
		MAIOR = B
		if A >= C{
			INTER, MENOR = A, C
		}else{
			INTER, MENOR = C, A
		}
	}else{
		MAIOR = C
		if  A >= B{
			INTER, MENOR = A, B
		}else{
			INTER, MENOR = B, A
		}
	}

	fmt.Printf("%.2f %.2f %.2f\n", MENOR, INTER, MAIOR)
}