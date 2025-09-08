package main

import "fmt"

func main() {

	var name1, name2 = "go", "programming language"

	fmt.Printf("%s is a %s.\n", name1, name2)

	var num1, num2, num3 = 10, 15, 25

	fmt.Printf("%d+%d=%d\n", num1, num2, num3)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/fmt
$ go run printf-main.go 
go is a programming language.
10+15=25

