package main
import "fmt"

func main(){
	var domains [2]string
	fmt.Println("Current values for array:", domains)

	// add value and print if
	domains[0] = "First element"
	fmt.Println("Set value: ", domains)
	fmt.Println("Get value for 0 element : ", domains[0])

	// get array length
	fmt.Println("Array length: ", len(domains))
	fmt.Println("Current values for array:", domains)

	// add one more value
	domains[1] = "Second element"

	// get array length
	fmt.Println("Array length: ", len(domains))
	fmt.Println("Current values for array:", domains)

	// use for loop to print our array
	for i := 0; i < len(domains); i++ {
		fmt.Println("Get value for element ", i, " is ", domains[i])
	}
}
