package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Welcome to Golang Programming language"
	str2 := "This is the Tutorial of Golang"

	fmt.Println("String before trimming")
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)

	result1 := strings.TrimPrefix(str1, "Welcome to")
	result2 := strings.TrimPrefix(str2, "This is")

	fmt.Println("String before trimming")
	fmt.Println("str1:", result1)
	fmt.Println("str2:", result2)
}
-----------------------

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Strings
$ go run trim-prefix.go 
String before trimming
str1: Welcome to Golang Programming language
str2: This is the Tutorial of Golang
String before trimming
str1:  Golang Programming language
str2:  the Tutorial of Golang
