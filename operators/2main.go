package main

import "fmt"

func main() {

	var a int = 30
	var b int = 50

	a %= b

	fmt.Println("A = ", a)

}


o/p ->
Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/operators
$ go run main.go
A =  30
