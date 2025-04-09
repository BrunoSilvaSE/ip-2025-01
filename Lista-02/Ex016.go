package main

import(
	"fmt"
	"math"
)

func main(){
	var A, B, C, delta float64

	fmt.Println("Digitie em sequencia os valores de A, B e C")
	fmt.Scan(&A, &B, &C)

	delta = math.Pow(B, 2) - 4*A*C

	if delta == 0{
		x := -B/2*A
		fmt.Printf("RAIZ ÚNICA\nf(x) = %.2fx² + %.2fX + %.2f\nX = %.2f\n", A, B, C, x)
	}else if delta > 0{
		x1 := (-B+math.Sqrt(delta))/float64(2*A)
		x2 := (-B-math.Sqrt(delta))/float64(2*A)
		fmt.Printf("RAÍZES DISTINTAS\nf(x) = %.2fx² + %.2fX + %.2f\nX = %.2f\nX` = %.2f\n", A, B, C, x1, x2)
	}else if delta < 0{
		delta = delta * -1
		x1 := fmt.Sprintf("%.2f + (%.2f)", -B/2*A, +math.Sqrt(delta)/2*A)
		x2 := fmt.Sprintf("%.2f + (%.2f)", -B/2*A, -math.Sqrt(delta)/2*A)
		fmt.Printf("RAÍZES IMAGINÁRIAS\nf(x) = %.2fx² + %.2fX + %.2f\nX = %s\nX` = %s\n", A, B, C, x1, x2)
	}

}