package main
import "fmt"

// function that accept int values and returns int
func total(x int, y int) int {
	return x + y
}

func main(){
	answer := total(10, 20)
	fmt.Println("10 + 20 = ", answer)
	fmt.Println("100 + 500 = ", total(100, 500))
}