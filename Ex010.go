package main

import(
	"fmt"
)

func main(){
	var a, b, c, d, det float32

	fmt.Println("Digite em sequencia os quatro elementos de formam uma matriz")
	fmt.Scan(&a ,&b ,&c ,&d)

	det = a*d-b*c

	fmt.Printf("O valor do determinante é = %.2f\n", det)
}