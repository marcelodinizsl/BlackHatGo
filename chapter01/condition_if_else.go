package main
import "fmt"

func main(){
	var number int
	fmt.Print("Enter your number: ")
	fmt.Scanf("%d", &number)
	fmt.Println(number)

	if number >= 0{
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is negative")
	}
}