package main

import "fmt"

func main() {

	var p *int

	if p == nil {

		fmt.Printf("the value of P is %x\n", p)
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/pointers
$ go run 2main.go
the value of P is 0
