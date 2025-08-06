
## complex number
package main

import "fmt"

func main() {

	var a = 3 + 5i
	var b = 2 + 3i

	var result1 = a + b
	var result2 = b - a
	var result3 = a * b
	var result4 = a / b

	fmt.Println(result1, result2, result3, result4)

}

########3
Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/datatypes
$ go run main.go
(5+8i) (-1-2i) (-9+19i) (1.6153846153846154+0.07692307692307686i)
