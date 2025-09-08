package main

import "fmt"

func main() {

	name1 := "kim"

	var name2 string = "Go Programming \n language"

	var name3 string = `Go programming \n langauage`

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)

	
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Strings
$ go run main.go 
kim
Go Programming 
 language
Go programming \n langauage


