package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str := "Golang Programming language"

	fmt.Println("the value of str1 is:", str)

	length := len(str)

	fmt.Println("the value of str is :", length)

	length2 := utf8.RuneCountInString(str)

	fmt.Println("the value of length2 with RuneCountInString is :", length2)
}
Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Strings
$ go run length-string.go 
the value of str1 is: Golang Programming language
the value of str is : 27
the value of length2 with RuneCountInString is : 27
