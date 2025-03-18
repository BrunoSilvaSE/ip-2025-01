// 2 Arrecadação de Jogos
package main

import (
	"fmt"
	"strings"
)

func main(){
	var nTeste int
	jogos := make(map[int]float32) 

	//Título
	fmt.Println(strings.Repeat("=", 120), "\n Programa de Arrecadação de Jogos\n", strings.Repeat("=", 120))

	// Input dos dados
	fmt.Println("Qual o número de casos de teste que deseja realiza?")

	fmt.Scan(&nTeste)

	fmt.Println(strings.Repeat("=", 120))
	fmt.Printf("%sAgora informe em sequência\n\n",strings.Repeat(" ", 46))
	fmt.Println("• O número de pessoas que compraram ingresso para o jogo correspondente ao caso de teste.")
	fmt.Println("• A percentagem de pessoas que compraram ingresso na categoria Popular.")
	fmt.Println("• A percentagem de pessoas que compraram ingresso na categoria Geral.")
	fmt.Println("• A percentagem de pessoas que compraram ingresso na categoria Arquibancada.")
	fmt.Println("• A percentagem de pessoas que compraram ingresso na categoria Cadeiras.")
	fmt.Println(strings.Repeat("=", 120))

	for i := 1; i <= nTeste; i++{
		var nPessoas int
		var pPopular, pGeral, pArquibancada, pCadeiras float32

		fmt.Printf("Digite as informações do %d° teste\n", i)
		fmt.Scan(&nPessoas, &pPopular, &pGeral, &pArquibancada, &pCadeiras)

		Arrecadacao := float32(nPessoas)*(pPopular + 5*pGeral + 10*pArquibancada + 20*pCadeiras)/100

		jogos[i] = Arrecadacao
	}

	//Print dos resultados
	for i, resultado := range jogos{
		fmt.Printf("a renda do jogo N. %d é = %.2f\n", i, resultado)
	}
}