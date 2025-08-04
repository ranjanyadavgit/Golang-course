package main

import "fmt"

func main() {

	var1 := 23
	var2 := "my first programming"
	var3 := 12.30

	fmt.Printf("the value of var 1 is: %d \n", var1)
	fmt.Printf("the type of var1 is: %T\n", var1)

	fmt.Printf("the value of the var2 is: %s\n ", var2)
	fmt.Printf("the type of the var2 is: %T\n", var2)

	fmt.Printf("the value of the var2 is: %f \n", var3)
	fmt.Printf("the type of the var2 is: %T\n", var3)

}

o/p ->

$ go run main.go
the value of var 1 is: 23 
the type of var1 is: int
the value of the var2 is: my first programming
 the type of the var2 is: string
the value of the var2 is: 12.300000
the type of the var2 is: float64
