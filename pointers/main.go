package main

import "fmt"

func main() {

	var p *int
	var a int = 100
	p = &a

	fmt.Printf("the value of a address: %x\n", &a)
	fmt.Printf("the address stored in p: %x\n", p)
	fmt.Print("the value of p is:", *p)

}



Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/pointers
$ go run main.go
the value of a address: c00000a0e8
the address stored in p: c00000a0e8
the value of p is:100
