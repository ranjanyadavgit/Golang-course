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

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/pointers
$ go run 4main.go
the value of x before calling x is 100
the value of x after calling x is 500

