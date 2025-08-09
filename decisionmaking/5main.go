package main

import "fmt"

func main() {

	var today int = 3

	switch {

	case today == 1:
		fmt.Println("today is sunday")

	case today == 2:
		fmt.Println("today is monday")
	case today == 3:
		fmt.Println("today is tuesday")

	case today == 4:
		fmt.Println("today is wednesday")
	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/decisionmaking
$  go run 5main.go
today is tuesday
