package main
import "fmt"

func main(){
	for i:=1; i <=10; i++{
		fmt.Printf("Welcome %d times.\n", i)

		// Break out of for loop when i is 5
		if i == 5 {
			break
		}
	}
}