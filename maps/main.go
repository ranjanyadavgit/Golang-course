package main

import "fmt"

func main() {

	var emptyMap = map[int]string{}

	fmt.Println("Empty Map:", emptyMap)

	var myMap = map[int]string{

		1: "ranjan",
		2: "kumar",
		3: "yadav",
		4: "bangalore",
		5: "electronic city",
	}
	fmt.Println("value of myMap:", myMap)
	fmt.Println("value of myMap:", myMap[2])
}



Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/maps
$ go run main.go
Empty Map: map[]
value of myMap: map[1:ranjan 2:kumar 3:yadav 4:bangalore 5:electronic city]
value of myMap: kumar
