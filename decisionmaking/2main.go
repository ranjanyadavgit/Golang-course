package main

import "fmt"

func main() {

	var a int = 20
	var b int = 15

	if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is greater than b")
	}

}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/decisionmaking
$ go run 2main.go 
a is greater than b
