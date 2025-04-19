package main

import "fmt"

func main(){
	var sum float64
	secI, sec := 1.0, 1.0

	for i := 1; i <= 50; i++{
		secI += 2
		sec++

		sum += secI/sec
	}

	fmt.Println(sum)
}