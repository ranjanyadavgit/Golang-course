package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "This is a go Programming language"
	str2 := "Welcome to Golang Programing Tutorial"

	fmt.Println("String before trimming")
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)

	result1 := strings.TrimSuffix(str1, "Programming language")
	result2 := strings.TrimSuffix(str2, "Tutorial")

	fmt.Println("String before trimming")
	fmt.Println("str1:", result1)
	fmt.Println("str2:", result2)
}
----------------

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Strings
$ go run trim-suffix.go 
String before trimming
str1: This is a go Programming language
str2: Welcome to Golang Programing Tutorial
String before trimming
str1: This is a go 
str2: Welcome to Golang Programing 
