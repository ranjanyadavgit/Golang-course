package main

import "fmt"

type address struct {
	state   string
	city    string
	zipcode int
}

func main() {

	a2 := address{state: "Bangalore", city: "Electronic", zipcode: 560001}

	fmt.Println("the addess:", a2)

	a3 := address{state: "jamshedpur"}

	fmt.Println("the address is", a3)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/struct
$ go run 1main.go 
the addess: {Bangalore Electronic 560001}
the address is {jamshedpur  0}
