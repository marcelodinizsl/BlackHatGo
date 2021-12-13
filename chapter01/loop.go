package main

import (
	"fmt"
)

func main(){
	m := 1
	for m <= 5 {
		fmt.Printf("Welcome %d times.\n", m)
		m = m + 1
	}

	for i:= m; i <= 10; i++ {
		fmt.Printf("Welcome %d times.\n", i)
	}
}