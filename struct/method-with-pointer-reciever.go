package main

import "fmt"

type author struct {
	name   string
	branch string
	salary int
}

func (a *author) show(abranch string) {
	(*a).name = abranch
}

func main() {

	result := author{
		name:   "Ranjan",
		branch: "CSE",
	}

	fmt.Println("before changing")
	fmt.Println("Author's name:", result.name)
	fmt.Println("Author's Branch", result.branch)

	p := &result
	p.show("ECE")
	fmt.Println("After changing")
	fmt.Println("Author's name", result.name)
	fmt.Println("Author's Branch", result.branch)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/struct
$ go run 5main.go 
before changing
Author's name: Ranjan
Author's Branch CSE
After changing
Author's name ECE
Author's Branch CSE
