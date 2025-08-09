package main

import "fmt"

func main() {

	var a int = 300
	var b int = 500

	if a == 300 {

		if b == 500 {

			fmt.Println("Value of A and B are Equal")

		}

	}
	fmt.Printf("Value of A: %d\n", a)
	fmt.Printf("Value of B: %d\n", b)
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/decisionmaking
$ go run 3main.go
Value of A and B are Equal
Value of A: 300
Value of B: 500
