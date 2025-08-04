package main

import "fmt"

func main() {

	var1, var2, var3 := 23, 24, 25

	fmt.Printf("the value of var 1 is: %d \n", var1)
	fmt.Printf("the type of var1 is: %T\n", var1)

	fmt.Printf("the value of the var2 is: %d\n ", var2)
	fmt.Printf("the type of the var2 is: %T\n", var2)

	fmt.Printf("the value of the var2 is: %d \n", var3)
	fmt.Printf("the type of the var2 is: %T\n", var3)

}



Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/variables
$ go run main.go
the value of var 1 is: 23 
the type of var1 is: int
the value of the var2 is: 24
 the type of the var2 is: int
the value of the var2 is: 25 
the type of the var2 is: int
