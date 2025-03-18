package main

import(
	"fmt"
)

func main(){
	var n int

	fmt.Println("Digite o número que sera validado.")
	fmt.Scan(&n)

	if n % 5 == 0 &&  n % 3 == 0 {
		fmt.Println("O NUMERO É DIVISIVEL")
	}else{
		fmt.Println("O NUMERO NÃO É DIVISIVEL")
	}
}