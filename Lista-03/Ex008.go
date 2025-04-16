package main

import "fmt"

func main(){
	var sum int
	
	for i := 1; i <= 20; i++{
		fmt.Printf("%d ", i)
		sum += i
	}

	fmt.Printf("\nsoma: %d\n", sum)
}