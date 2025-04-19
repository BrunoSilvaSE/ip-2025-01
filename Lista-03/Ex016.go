package main

import "fmt"

func main(){
	var n1, n2, n3, Nqtd int
	sec := []int{}

	fmt.Println("Digite dois valores inteiros e positivos")
	fmt.Scan(&n1, &n2)

	sec = append(sec, n1)
	sec = append(sec, n2)

	fmt.Println("Quantos termos deseja exibir?")
	fmt.Scan(&Nqtd)

	for i := 1; i <= Nqtd-2; i++{
		if i % 2 != 0{// O núemro é impar
			n3 = n1 + n2
		}else{ // O número é par
			n3 = n1 - n2
		}
		n1 = n2
		n2 = n3
		sec = append(sec, n3)
	}

	fmt.Println("Série de Fetuccine")
	fmt.Println(sec)
}