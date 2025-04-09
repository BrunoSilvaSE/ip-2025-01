package main

import (
	"fmt"
)

func main(){
	//Declaração
	var reg, tipe int
	
	tab := map[int][3]string{// "destino" : {ida, ida e volta}
		1 : {"1 - Região Norte", "500", "900"},
		2 : {"2 - Região Nordeste", "350", "650"},
		3 : {"3 - Região Centro-Oeste", "350", "600"},
		4 : {"4 - Região Sul", "300", "550"},
	}

	// input região
	fmt.Printf("Digite o valor que faz referência ao local de destino\n1 - Região Norte\n2 - Região Nordeste\n3 - Região Centro-Oeste\n4 - Região Sul\n",)
	fmt.Scan(&reg)
	if reg > 4 || reg < 0 {
		fmt.Println("O valor digitado é invalido!")
		return
	}

	// input tipo de vôo
	fmt.Printf("A viagem inclui retorno\n[1] Sim, Ida e Volta\n[2] Não, só Ida\n",)
	fmt.Scan(&tipe)
	if tipe > 2 || tipe < 1 {
		fmt.Println("O valor digitado é invalido!")
		return
	}

	//retorno
	if tipe == 1{
		fmt.Printf("Destino: %s\nTipo: Ida e Volta\nValor: %s", tab[reg][0], tab[reg][2])
	}else{
		fmt.Printf("Destino: %s\nTipo: Ida\nValor: %s\n", tab[reg][0], tab[reg][2])
	}
}