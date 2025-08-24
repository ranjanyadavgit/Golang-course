package main

import "fmt"

func main() {

	func(name string) {

		fmt.Println(name)

	}("Welcome to Go programming language")

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/functions
$ go run 10main.go 
Welcome to Go programming language
