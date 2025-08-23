package main

import "fmt"

func main() {

	var myMap = map[int]string{

		10: "ranjan",
		20: "kim",
		30: "alex",
		40: "mark",
	}

	for id, name := range myMap {

		fmt.Printf("ID:%d\t%s\n", id, name)
	}
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/maps
$ go run 6main.go
ID:10   ranjan
ID:20   kim
ID:30   alex
ID:40   mark
