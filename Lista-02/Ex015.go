package main

import(
	"fmt"
)

func main(){
	var D, M, A int
	meses := map[int]string{
		1: "Janeiro", 2: "Fevereiro", 3: "Março",
		4: "Abril", 5: "Maio", 6: "Junho",
		7: "Julho", 8: "Agosto", 9: "Setembro",
		10: "Outubro", 11: "Novembro", 12: "Dezembro",
	}

	fmt.Println("Digite em sequencia o dia, mês e ano")
	fmt.Scan(&D, &M, &A)
	fmt.Printf("%d de %s de %d\n", D, meses[M], A)
}