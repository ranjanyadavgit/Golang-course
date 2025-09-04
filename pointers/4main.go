package main

import "fmt"

func display(a *int) {

	*a = 500
}
func main() {

	var x = 100

	fmt.Printf("the value of x before calling x is %d\n", x)
	display(&x)
	fmt.Printf("the value of x after calling x is %d\n", x)
}
