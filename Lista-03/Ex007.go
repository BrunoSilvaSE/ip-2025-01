package main

import(
	"fmt"
)

func main(){
	var n, sumN, sumNP, maior, qtdN, qtdNP, qtdNI int 
	menor := 9999999999
	var medN, medNP, percNI float64

	for {
		fmt.Println("Digite um número: (caso queire parar digite 30000)")
		fmt.Scan(&n)
		if n == 30000{break}

		//a) a soma dos números digitados;
		sumN += n
		//b) a quantidade de números digitados;
		qtdN++
		//c) a média dos números digitados;
		//depois do for
		//d) o maior número digitado;
		if maior < n {maior = n}
		//e) o menor número digitado;
		if menor > n {menor = n}
		//f) a média dos números pares;
		if n % 2 == 0 {
			sumNP += n
			qtdNP++
		}
		//g) a percentagem dos números ímpares entre todos os números digitados.
		if n % 2 != 0 { qtdNI++ }
	}

	//c) a média dos números digitados;
	if qtdN != 0 {
		medN = float64(sumN)/float64(qtdN)
	}
	//f) a média dos números pares;
	if qtdNP != 0 {
		medNP = float64(sumNP)/float64(qtdNP)
	}
	//g) a percentagem dos números ímpares entre todos os números digitados.
	if qtdN != 0{
		percNI = (float64(qtdNI)/float64(qtdN))*100
	}

	fmt.Println("a) a soma dos números digitados; ", sumN)
	fmt.Println("b) a quantidade de números digitados; ", qtdN)
	fmt.Println("c) a média dos números digitados; ", medN)
	fmt.Println("d) o maior número digitado; ", maior)
	fmt.Println("e) o menor número digitado; ", menor)
	fmt.Println("f) a média dos números pares; ", medNP)
	fmt.Println("g) a percentagem dos números ímpares entre todos os números digitados. ", percNI)
}