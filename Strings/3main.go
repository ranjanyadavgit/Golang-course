package main

import "fmt"

func main() {

	str1 := "Go Programming language"

	for index, ch := range str1 {

		fmt.Printf("the Index is %d and character is %c\n", index, ch)
	}
}
-------

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/Strings
$ go run 3main.go 
the Index is 0 and character is G
the Index is 1 and character is o
the Index is 2 and character is  
the Index is 3 and character is P
the Index is 4 and character is r
the Index is 5 and character is o
the Index is 6 and character is g
the Index is 7 and character is r
the Index is 8 and character is a
the Index is 9 and character is m
the Index is 10 and character is m
the Index is 11 and character is i
the Index is 12 and character is n
the Index is 13 and character is g
the Index is 14 and character is
the Index is 15 and character is l
the Index is 16 and character is a
the Index is 17 and character is n
the Index is 18 and character is g
the Index is 19 and character is u
the Index is 20 and character is a
the Index is 21 and character is g
the Index is 22 and character is e

