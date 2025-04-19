package main

import "fmt"

func main(){
	var n, termo int
	termoOLD := 0
	sec := []int{}
	r := 1

	fmt.Println("Digite o numero de termos desejados:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		termo = termoOLD + r
		termoOLD = termo
		sec = append(sec, termo)
		r += 2
	}


	fmt.Println("Sequência:")
	fmt.Println(sec)
}