package main

import "fmt"

var globalVariable int = 100

func main() {

	var localVariable int = 200

	fmt.Printf("the value of the global variable is: %d\n", globalVariable)
	fmt.Printf("the value of the local variable is: %d\n", localVariable)

	display()

}

func display() {

	fmt.Printf("the value of the global variable is: %d\n", globalVariable)
}
