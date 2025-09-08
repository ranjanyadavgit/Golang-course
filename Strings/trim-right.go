package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "!!Welcome to Go programming Language **"

	str2 := "@@ this is a golang tutorial $$"

	fmt.Println("String before trimming")
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)

	result1 := strings.TrimRight(str1, "*")
	result2 := strings.TrimRight(str2, "$")

	fmt.Println("String before trimming")
	fmt.Println("str1:", result1)
	fmt.Println("str2:", result2)

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Strings
$ go run trim-left.go 
String before trimming
str1: !!Welcome to Go programming Language **
str2: @@ this is a golang tutorial $$
String before trimming
str1: !!Welcome to Go programming Language
str2: @@ this is a golang tutorial
