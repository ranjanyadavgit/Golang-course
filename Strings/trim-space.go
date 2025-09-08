package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "  **Welcome to Go programming langaage** "
	str2 := "   @@ this is golang programming language@@   "

	fmt.Println("String before trimming")
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)

	result1 := strings.TrimSpace(str1)
	result2 := strings.TrimSpace(str2)

	fmt.Println("String after trimming")
	fmt.Println("str1:", result1)
	fmt.Println("str2", result2)

}

-------------
Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Strings
$ go run trimspace.go 
String before trimming
str1:   **Welcome to Go programming langaage** 
str2:    @@ this is golang programming language@@   
String after trimming
str1: **Welcome to Go programming langaage**
str2 @@ this is golang programming language@@
