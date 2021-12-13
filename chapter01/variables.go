package main

import(
	"fmt"
)

func main(){
	const pi float32 = 3.14159
	var i int
	i = 10

	var y = 5
	var msg = "Remote host found"
	var foo int = 100

	vehicle := "Mercedes"
	age := 52

	fmt.Printf("%d %d %s\n", i, y, msg)
	fmt.Printf("%d", foo)
	fmt.Printf("%d", age)
	fmt.Printf("%s", vehicle)
	fmt.Printf("%f", pi)
}
