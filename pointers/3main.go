package main

import "fmt"

func main() {

	var a = 450
	var p = &a

	fmt.Println("the value of a is:", a)
	fmt.Println("the address stored in a is :", &a)
	fmt.Println("the value of p is:", p)
	fmt.Println("the value of p before changing is:", *p)

	*p = 500
	fmt.Println("value stored in a after changing:", a)
	
Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/pointers
$ go run 3main.go
the value of a is: 450
the address stored in a is : 0xc00000a0e8
the value of p is: 0xc00000a0e8
the value of p before changing is: 450
value stored in a after changing: 500
