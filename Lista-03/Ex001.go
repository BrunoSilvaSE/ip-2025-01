package main

import(
	"fmt"
)

func main(){
	var expo int
	var n float64
	poten := 1.00

	fmt.Println("Qual o valor da base da potência?")
	fmt.Scan(&n)

	for true {
		fmt.Println("Qual o expoente? lembre-se que deve ser um valor inteiro maior que -1")
		fmt.Scan(&expo)

		if expo < 0 {
			fmt.Println("O valor digitado é inváldio, por favor tente novamente!")
		}else{
			break
		}
	}

	if expo == 0 && n == 0{
		err := "O valor da potência é indeterminado"
		fmt.Println(err)
		return
	}

	for i := 1; i <= expo; i++{
		poten = poten * n
	}

	fmt.Printf("%.2f elevado á %d é igual a %.2f\n", n, expo, poten)

}