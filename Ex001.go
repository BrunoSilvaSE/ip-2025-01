//1 Aprovado ou Reprovado
package main

import (
	"fmt"
)

func main()  {
	var n1, n2, n3 float32
	var med float32
	
	fmt.Println("Digite as 3 notas para calcular a média")
	fmt.Scan(&n1, &n2, &n3)

	med = (n1+n2+n3)/3
	
	if med > 6 {
		fmt.Printf("Parabéns!!, você foi APROVADO e sua média final foi de: %.2f" , med)
	}else{
		fmt.Printf("Você foi REPROVADO, sua média final foi de: %.2f" , med)
	}
}