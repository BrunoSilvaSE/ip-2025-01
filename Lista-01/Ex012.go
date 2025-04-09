package main

import(
	"fmt"
)

func main(){
	var hora float32

	fmt.Println("Digite o número de horas em que a charrete foi usada.")
	fmt.Scan(&hora)

	n3 := float32((hora - float32(int(hora)))*5) // multiplica os valores apos as vírgulas por 5
	n1 := int(hora/3)*10 // multiplica os blocos de 3h por 10
	n2 := float32((int(hora)%3)*5) // multiplica o resto das horas que não agrupam em 3 por 5
	fmt.Println("O valor a pagar é = R$",float32(n1)+n2+n3)
}