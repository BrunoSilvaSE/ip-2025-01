package main

import (
	"fmt"
	"math"
)

func main(){
	var n int

	for {
		fmt.Scan(&n)
		if n <= 0{break}

		qrt := math.Pow(float64(n), 0.5)
		resto := qrt - float64(int(qrt))

		if resto == 0{
			fmt.Printf("%d é um quadrado perfeito!\n", n)
		}else{
			fmt.Printf("%d não é um quadrado perfeito!\n", n)
		}
	}
}