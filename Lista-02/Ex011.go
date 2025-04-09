package main

import(
	"fmt"
)

func main(){
	var n float64

	fmt.Println("Digite um valor númerico")
	fmt.Scan(&n)

	fmt.Printf("Dado f(x) = 8/(2-x), f(%.2f) = %.2f\n", n, 8/(2-n))
}