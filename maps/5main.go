package main

import "fmt"

func main() {

	myMap := map[int]string{

		10: "Ranjan",
		20: "Yadav",
		30: "Bangalore",
		40: "electronic-city",
		50: "Tcs",
	}

	fmt.Println("mymap value is", myMap)

	myMap[60] = "DevOps"
	myMap[70] = "DevSecops"

	fmt.Println("After adding new elements", myMap)

	myMap[20] = "Kumar"
	myMap[30] = "Kolkata"

	fmt.Println("after updating the values", myMap)

	delete(myMap, 50)

	fmt.Println("After deleting the 50 index values", myMap)

}


Lenovo@DESKTOP-M5DT73G MINGW64 /d/Study/Golang/src/maps
$ go run 5main.go
mymap value is map[10:Ranjan 20:Yadav 30:Bangalore 40:electronic-city 50:Tcs]
After adding new elements map[10:Ranjan 20:Yadav 30:Bangalore 40:electronic-city 50:Tcs 60:DevOps 70:DevSecops]
after updating the values map[10:Ranjan 20:Kumar 30:Kolkata 40:electronic-city 50:Tcs 60:DevOps 70:DevSecops]
After deleting the 50 index values map[10:Ranjan 20:Kumar 30:Kolkata 40:electronic-city 60:DevOps 70:DevSecops]
