package main

import "fmt"

func main() {

	a := [...]int{1, 3, 4, 5}
	names := [4]string{"ranjan", "kumar", "yadav", "hey"}
	fmt.Println("length of array", len(a))
	fmt.Printf("name: %v\n", names)

	fmt.Println(names[:])
	fmt.Println(names[:3])

	fmt.Println(names[2:])
	fmt.Println(a[1:3])
	fmt.Println(names[1:3])
}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 4main.go
length of array 4
name: [ranjan kumar yadav hey]
[ranjan kumar yadav hey]
[ranjan kumar yadav]
[yadav hey]
[3 4]
[kumar yadav]

Lenovo@DESKTOP-M5DT73G MINGW64 /d/S
