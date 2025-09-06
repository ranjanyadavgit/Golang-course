package main

import "fmt"

// Author struct
type Author struct {
	name   string
	branch string
	year   int
}

type HR struct {
	details Author
}

func main() {

	result := HR{
		details: Author{"Julia", "ECE", 2021},
	}
	fmt.Println("Details of Authors")
	fmt.Println(result)
}

Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/struct
$ go run 2main.go
Details of Authors
{{Julia ECE 2021}}
