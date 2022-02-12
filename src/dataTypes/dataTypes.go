package dataTypes

import (
	// formatting strings and console prints
	"fmt"
	// mathematical operations
	"math"
)

// Variables
var varString string = "string"

func varStringFunc() {
	// or shorthand, type is defined by assignement, can only be used inside of a function
	varString1 := "string"
	fmt.Println("\nfunc shorthand string: ", varString1)
}

var varInt int = 69

func varIntFunc(n int) {
	fmt.Println("\nfunc int: ", n)
}

// other numeric types https://golang.org/ref/spec#Numeric_types

// input type must be float64 and return type must be float64
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// Pointers
// function cannot access values of other functions or other variables outside function
// for that we use pointers that point to variable's address in memory
func updateWithPointer(x *string) {
	*x = "James"
}

func updateWithoutPointer(x string) {
	x = "John"
}

func pointers() {
	name := "Jack"
	memoryName := &name
	fmt.Println("\noriginal value:", name)
	fmt.Println("memory address:", memoryName)
	fmt.Println("value at memory address:", *memoryName)
	updateWithPointer(memoryName)
	fmt.Println("\nupdated value with pointer:", name)
	updateWithoutPointer(name)
	fmt.Println("\nupdated value without pointer:", name)
}

// Prints
func DataTypes() {
	fmt.Println("\nData types: ")
	varStringFunc()
	varIntFunc(69)
	pointers()

	// Printf( formatted strings) %_ = format specifier
	fmt.Printf("\nint is %v and string is %v", varInt, varString)
	fmt.Printf("\nint is %b and string is %q \n", varInt, varString)
	// other format specifiers https://pkg.go.dev/fmt

	// Sprintf save formatted strings
	var savedString = fmt.Sprintf("\nsaved string: \n	int is %v and string is %v \n", varInt, varString)
	fmt.Println(savedString)

	fmt.Println(circleArea(6.9))
}
