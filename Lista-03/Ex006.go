package main

import "fmt"

func main(){
	var n int
	prod := 1
	i := 1
	stt := false

	fmt.Scan(&n)

	for prod < n{
		prod = i * (i+1) * (i+2)
		i++
		if prod == n {
			stt = true
			break
		}
	}

	if stt{
		fmt.Printf("%d eh triangular\n", n)
	}else{
		fmt.Printf("%d não eh triangular\n", n)
	}
}