package main

import "fmt"

func main() {

	var emptyMap = map[int]string{}

	fmt.Println("empty string:", emptyMap)

	var value = map[int]string{

		1: "ranjan",
		2: "kumar",
		3: "yadav",
		4: "bangalore",
		5: "electronic city",
	}
	fmt.Println("value of map:", value)
	fmt.Println("value of map:", value[2])
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/maps
$ go run main.go
empty string: map[]
value of map: map[1:ranjan 2:kumar 3:yadav 4:bangalore 5:electronic city]
value of map: kumar
