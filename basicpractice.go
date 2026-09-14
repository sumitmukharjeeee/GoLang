package main

import "fmt"

func main() {
	var name string = "Sumit Mukharjee"
	var age int = 24

	fmt.Println(name)
	fmt.Println(age)

	var resultAge bool = isAdult(age)
	fmt.Println(resultAge)
}

func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

// two same name go functions cant exist in same folder or package I guess
