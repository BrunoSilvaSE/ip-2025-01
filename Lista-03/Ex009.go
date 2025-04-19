package main

import (
	"fmt"
)

func main(){
	var qtdAlunos, notasQtd, n1, n2 int
	var med, notasSum float64
	qtds := [3]int{0, 0, 0} // {qtdReprovado, qtdExame, qtdAprovado}

	fmt.Scan(&qtdAlunos)

	for i := 1; i <= qtdAlunos; i++{
		fmt.Scan(&n1, &n2)

		med = (float64(n1)+float64(n2))/2
		notasSum += med
		notasQtd++

		if med <= 3{
			qtds[0]++
		}else if med > 3 && med < 7{
			qtds[1]++
		}else if med >= 7{
			qtds[2]++
		}

		fmt.Printf("Média Aluno %d: %.2f\n", i, med)
	}

	medClass := notasSum/float64(notasQtd)

	fmt.Println("Total Aprovados: ", qtds[2])
	fmt.Println("Total Exame: ", qtds[1])
	fmt.Println("Total Reprovados: ", qtds[0])
	fmt.Println("Média da classe: ", medClass)

}