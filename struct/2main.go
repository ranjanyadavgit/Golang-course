package main

import "fmt"

type address struct {
	state   string
	city    string
	zipcode int
}

func main() {

	a4 := new(address)
	a4.state = "Karnataka"
	a4.city = "bangalore"
	a4.zipcode = 560010

	fmt.Println("the address:", a4)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/struct
$ go run 1main.go 
the address: &{Karnataka bangalore 560010}
