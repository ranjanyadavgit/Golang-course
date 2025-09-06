package main

import "fmt"

type address struct {
	state   string
	city    string
	zipcode int
}

func main() {

	var a5 = &address{}
	a5.state = "New York"
	a5.city = "New York City"
	a5.zipcode = 12345

	fmt.Println("the Address5:", a5)
}

$ go run 1main.go 
the Address5: &{New York New York City 12345}
