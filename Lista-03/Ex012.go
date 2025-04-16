package main

import (
	"fmt"
)

func main(){
	var n, sum, s float64

	fmt.Scan(&n)

	for i := 1; i <= 20; i++ {
		if i % 2 != 0{
			sum -= n/float64(fatorial(i))
		}else{
			sum += n/float64(fatorial(i))
		}
	}

	s = n + sum
	fmt.Println(s)
}

func fatorial(n int)int{
	fact := 1

	if n == 0 {return 1}

	for i := 1; i <= n; i++{
		fact *= i
	}

	return fact
}