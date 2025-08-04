package main

import "fmt"

func main() {

	var variable1, variable2, variable3 = 1, "Golang", 65.55

	fmt.Printf("the value of the variable1 is: %d\n", variable1)
	fmt.Printf("the type of the variable2 is: %T\n", variable1)

	fmt.Printf("the value of the variable2 is: %s\n", variable2)
	fmt.Printf("the type of the variable2 is: %T\n", variable2)

	fmt.Printf("the value of the variable3 is: %f\n", variable3)
	fmt.Printf("the type of the variable3 is: %T\n", variable3)
}
