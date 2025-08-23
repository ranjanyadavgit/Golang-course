package main

import "fmt"

func main() {

	var myMap = make(map[string]float64)

	myMap["ranjan"] = 60.44
	myMap["jim"] = 55.44
	myMap["alex"] = 65.22

	fmt.Println("the values are:", myMap)

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/maps
$ go run 2main.go
the values are: map[alex:65.22 jim:55.44 ranjan:60.44]
