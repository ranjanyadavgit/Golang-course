package main

import "fmt"

func main() {

	var result = "Go programming language"

	display(result)

	var result2 = 123
	display(result2)
}

func display(a interface{}) {

	value, ok := a.(string)
	fmt.Printf("Value:%v --- status is %v\n", value, ok)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Interfaces
$ go run type-assertionmain.go 
Value:Go programming language --- status is true
Value: --- status is false
