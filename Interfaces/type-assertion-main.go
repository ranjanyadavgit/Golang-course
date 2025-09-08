package main

import "fmt"

func main() {

	var result = "Go programming laanguage"
	display(result)
}

func display(a interface{}) {

	value := a.(string)
	fmt.Println("Value:", value)

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Interfaces
$ go run type-assertionmain.go 
Value: Go programming laanguage
