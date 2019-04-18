package main

import (
	"fmt"
)

// Defining struct data type using Person
type Employee struct {
	name, address string
	age           int
	salary        float64
	isTopper      bool
}

// Defining struct data type using Cat
type Cat struct {
	name string
}

// Defining struct data type using Dog
type Dog struct {
	name string
}

// Defining interface using Animal
type Animal interface {
	talking()
}

//Defining function with Cat receiver
func (c Cat) talking() {
	fmt.Println(c.name, " Saying Meow! Meow! ")
}

//Defining function with Dog receiver
func (d Dog) talking() {
	fmt.Println(d.name, " Saying Bow! Bow! ")
}

//Defining polymorphic method using interface
func Speaks(a Animal) {
	a.talking()
}

func main() {

	// Printing message using fmt package
	fmt.Println("Hello World!")

	// Initializing integer using short hand operator
	a := 2
	b := 3
	c := a + b
	fmt.Println(" Integers a = ", a, " b = ", b, " c = ", c)

	// Initializing floats using short hand operator
	a1 := 2.2
	b1 := 3.3
	c1 := a1 + b1
	fmt.Println(" Floats a1 = ", a1, " b1 = ", b1, " c1 = ", c1)

	// Initializing bool values
	IsValid := true
	IsWrong := false
	fmt.Println(" Bools IsValid = ", IsValid, " IsWrong = ", IsWrong)

	// Defining slices with ints
	IntAry := []int{45, 90, 99}
	fmt.Println("Int array = ", IntAry)

	// Defining slices with strings
	StringAry := []string{"Hyderabad", "Pune", "Mumbai"}
	fmt.Println("String array = ", StringAry)

	// Defining slices with multiple types
	MultiTypeAry := []interface{}{2, true, 89.96, false, "Hyderabad"}
	fmt.Println("Multi Type array = ", MultiTypeAry)

	//Defining slices with employee struct
	StructAry := []Employee{
		Employee{
			name:     "Jack",
			address:  "Ireland",
			age:      32,
			salary:   98987.22,
			isTopper: true,
		},
		Employee{
			name:     "Steffan",
			address:  "USA",
			age:      88,
			salary:   8989.22,
			isTopper: false,
		},
	}
	fmt.Println("Struct Array ", StructAry)

	//Defining map with marks
	Marks := map[string]int{
		"English": 99,
		"Hindi":   98,
		"Maths":   100,
		"Social":  99,
	}
	fmt.Println("Marks map = ", Marks)

	//Defining  map with multiple types

	Person := map[string]interface{}{
		"name":             "Sateesh",
		"age":              32,
		"isGoogleEmployee": false,
		"salary":           7890.98,
	}
	fmt.Println("Person Details = ", Person)

	PersonsMapAry := []map[string]interface{}{
		map[string]interface{}{
			"name":             "Sai",
			"age":              32,
			"isGoogleEmployee": false,
			"salary":           7890.98,
		},
		map[string]interface{}{
			"name":             "Rama",
			"age":              99,
			"isGoogleEmployee": true,
			"salary":           7890989.98,
		},
	}

	fmt.Println("Persons Map Array = ", PersonsMapAry)

	// Defining Cat and Dog struct types.

	jimmy := Cat{
		name: "Jimmy",
	}

	puppy := Dog{
		name: "Puppy",
	}

	//We can listen to their sounds using interface.
	Speaks(jimmy)
	Speaks(puppy)

}
