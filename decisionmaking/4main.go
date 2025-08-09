package main

import "fmt"

func main() {

	var a int = 200
	var b int = 200

	if a < b {
		fmt.Println("a is lesser than b")
	} else if a == b {
		fmt.Println("a is equal to b")

	} else {
		fmt.Println("a is greater than b")
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/decisionmaking
$ go run 4main.go
a is equal to b

