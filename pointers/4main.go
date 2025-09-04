package main

import "fmt"

func display(a *int) {

	*a = 500
}
func main() {

	var x = 100

	fmt.Printf("the value of x before calling x is %d\n", x)
	display(&x)
	fmt.Printf("the value of x after calling x is %d\n", x)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/Golang_basics/pointers
$ go run 4main.go
the value of x before calling x is 100
the value of x after calling x is 500

-------
####
#This Go code demonstrates pass by reference using pointers. Let me break it down:

#Code Explanation
#Function display(a *int):

#Takes a parameter a which is a pointer to an integer (*int)
#*a = 500 dereferences the pointer and sets the value at that memory location to 500
#This modifies the original variable that the pointer points to

#Function main():

#var x = 100 declares and initializes variable x with value 100
#First Printf prints the initial value of x (100)
#display(&x) calls the function, passing the address of x using the & operator
#Second Printf prints the value of x after the function call

#Key Concepts
#Pointers in Go:

#*int means "pointer to an integer"
#&x gets the memory address of variable x
#*a dereferences the pointer to access/modify the actual value

#Pass by Reference:

#By passing &x (address of x), the function can modify the original variable
#Without pointers, Go passes copies of values, so changes wouldn't affect the original

#Output
#the value of x before calling x is 100
#the value of x after calling x is 500
