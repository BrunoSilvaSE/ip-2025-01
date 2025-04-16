package main

import(
	"fmt"
)

func main(){
	var sum, qtN, med int

	for n := 50; n <= 70; n++{
		if n % 2 == 0{
			sum += n
			qtN++
		}
	}

	med = sum/qtN

	fmt.Println("Soma e media dos valor de 50 a 70:")
	fmt.Printf("Soma: %d\n", sum)
	fmt.Printf("Média %.2f\n", med)
	println(qtN)
}