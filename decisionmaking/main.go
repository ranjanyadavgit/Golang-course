package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	if a < b {

		fmt.Println("a value is less than b value")
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/decisionmaking
$ go run main.go
a value is less than b value

