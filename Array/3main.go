package main

import "fmt"

func main() {

	x := [5]int{1: 10, 3: 30}
	y := [6]int{2: 5, 3: 6}
	z := [10]int{3: 7, 5: 7, 7: 8}
	k := [4]float64{2: 3.4, 3: 4.5}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(k)

}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/array
$ go run 3main.go
[0 10 0 30 0]
[0 0 5 6 0 0]
[0 0 0 7 0 7 0 8 0 0]
[0 0 3.4 4.5]

