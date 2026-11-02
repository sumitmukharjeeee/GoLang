// Why do we need syructs
// Imagine we want to represent a student we have

// name := "Sumit"
// age := 32
// course := "Golang"
// graduateYear := 2023

// but thee variables are unrelated as gos type system
// what if we have 100s of them we need something meaningful to group them right??
// this is where structs come in

package main

import "fmt"

// These are custome Go types
//Student1 os one instance of student
type Student struct {
	Name         string
	Age          int
	Course       string
	GraduateYear int
}

func main() {

	// Student is a type
	// Name age etc are fields of the type students
	// student1 is a variavles containing sturct values

	student1 := Student{
		Name:         "Rahul",
		Age:          32,
		Course:       "Golang",
		GraduateYear: 2023,
	}
	student2 := Student{
		Name:"Amit",
		// remaning fields will get zero values
		// Tipe - zero values are important and frequently used in Go
	}
	fmt.Println(student1)
	// Accessing and modyfing fields
	fmt.Println(student1.Name)
	student1.Age = 28
	fmt.Printf("Age: %d\n", student1.Age)
	fmt.Println(student2)
}
