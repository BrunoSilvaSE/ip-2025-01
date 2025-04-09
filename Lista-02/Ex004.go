package main

import(
	"fmt"
	"math"
)

func main(){
	var n float64

	fmt.Println("Digite um número: ")
	fmt.Scan(&n)

	if n < 0{
		fmt.Println(math.Pow(n,2))
	}else{
		fmt.Println(math.Sqrt(n))
	}
}