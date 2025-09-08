package main

import "fmt"

func main() {

	var i int = 100
	var s string = "Hello"

	kind(i)
	kind(s)
}

func kind(i interface{}) {
	fmt.Printf("the type of %v is %T\n", i, i)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Interfaces
$ go run emptyinterfacemain.go 
the type of 100 is int
the type of Hello is string
