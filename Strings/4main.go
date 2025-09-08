package main

import "fmt"

func main() {

	str := "Go Programming Language"

	for i := 0; i < len(str); i++ {

		fmt.Printf("index %d = %c and ASCII code = %v\n", i, str[i], str[i])

	}
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Strings
$ go run 4main.go 
index 0 = G and ASCII code = 71
index 1 = o and ASCII code = 111
index 2 =   and ASCII code = 32
index 3 = P and ASCII code = 80
index 4 = r and ASCII code = 114
index 5 = o and ASCII code = 111
index 6 = g and ASCII code = 103
index 7 = r and ASCII code = 114
index 8 = a and ASCII code = 97
index 9 = m and ASCII code = 109
index 10 = m and ASCII code = 109
index 11 = i and ASCII code = 105
index 12 = n and ASCII code = 110
index 13 = g and ASCII code = 103
index 14 =   and ASCII code = 32
index 15 = L and ASCII code = 76
index 16 = a and ASCII code = 97
index 17 = n and ASCII code = 110
index 18 = g and ASCII code = 103
index 19 = u and ASCII code = 117
index 20 = a and ASCII code = 97
index 21 = g and ASCII code = 103
index 22 = e and ASCII code = 101
