package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Welcome, to the, Golang programming, language"
	str2 := "My Name is ranjan"
	str3 := "hi there this is a programming"

	fmt.Println("before split")
	fmt.Println("the value:", str1)

	result1 := strings.Split(str1, ",")
	result2 := strings.Split(str2, "")
	result3 := strings.Split(str3, "!")

	fmt.Printf("\n Result 1 Type :%T \t Result:%v\n", result1, result1)

	for _, value := range result1 {

		fmt.Println(value)
	}

	fmt.Println("\nResult2:", result2)
	fmt.Println("result3:", result3)

}
===============
$ go run split.go 
before split
the value: Welcome, to the, Golang programming, language

 Result 1 Type :[]string         Result:[Welcome  to the  Golang programming  language]
Welcome
 to the
 Golang programming
 language

Result2: [M y   N a m e   i s   r a n j a n]
result3: [hi there this is a programming]
