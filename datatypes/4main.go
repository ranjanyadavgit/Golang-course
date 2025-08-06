#boolean
package main

import "fmt"

func main() {

	str1 := "go programming"
	str2 := "GO PROGRAMMING"
	str3 := "GO Programming"

	result1 := str1 == str2
	result2 := str1 == str3

	fmt.Println("result 1 : ", result1)
	fmt.Println("result2 : ", result2)

}

o/p - > Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/datatypes
$ go run main.go
result 1 :  false
result2 :  false
