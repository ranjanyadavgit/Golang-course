package main

import "fmt"

func main() {

	var year int = 5
	switch {

	case year == 1:
		fmt.Println("this is january")

	case year == 2:
		fmt.Println("this is february")

	case year == 3:
		fmt.Println("this is march")

	case year == 4:
		fmt.Println("this is april")

	case year == 5:
		fmt.Println("this is may")

	case year == 6:
		fmt.Println("this is june")

	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/decisionmaking
$  go run 6main.go
this is may
