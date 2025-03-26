//3 Composição Inteira (+)
package main

import(
	"fmt"
)

func main(){
	var c, d, u int

	//input dos valores
	fmt.Println("Digite em sequência 3 números inteiros correspondentes a casa das centenas, dezenas e unidade")
	fmt.Scan(&c, &d, &u)

	//processamento de dados
	c = c*100
	d = d*10
	num := u+d+c

	quad := num*num

	//verificando se foi digitado algum número de 2 dígitos ou mais
	if num > 999 || num < 0{
		fmt.Println("DIGITO INVALIDO\nDigite apenas números entre 0 e 9")
		return
	}

	//printando as informções
	fmt.Printf("O resultado da composição dos 3 núemros é: %d\n", num)
	fmt.Printf("e séu quadrado é: %d\n", quad)
}