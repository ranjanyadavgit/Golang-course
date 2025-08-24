package main

import "fmt"

func swap(x *int, y *int) {

	var temp int

	temp = *x
	*x = *y
	*y = temp
}

func main() {

	var a int = 100
	var b int = 200

	fmt.Printf("before swap, A = %d,B = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After Swap, A = %d,B = %d", a, b)

}
Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/functions
$ go run 7main.go
before swap, A = 100,B = 200
After Swap, A = 200,B = 100
