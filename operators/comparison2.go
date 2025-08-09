package main

import "fmt"

func main() {

	a := 50
	b := 40

	result1 := a >= b

	fmt.Println("A >= B :", result1)
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/operators
$ go run comparison.go 
A >= B : true
