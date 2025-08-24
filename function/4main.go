package main

import (
	"fmt"
	"math"
)

func main() {

	square := func(x float64) float64 {

		return math.Sqrt(x)
	}

	fmt.Println(square(16))
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/functions
$ go run 4main.go
4
