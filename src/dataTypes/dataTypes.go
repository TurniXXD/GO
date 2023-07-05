package dataTypes

import (
	// formatting strings and console prints
	"bytes"
	"fmt"

	// mathematical operations
	"math"
)

// Variables
var varString string = "string"

func varStringFunc() {
	// One char in string is byte that uses binary representation of decimal number in ASCII
	// Otherwise string is grapheme decimal number that uses standard unicode
	// In unicode we can also combine codes like code 128077 👍 + 127999 🟫 generates 👍🏿, this is called emoji modifier
	// We then convert graphemes to binary, this operation is called unicode encoding
	// or shorthand, type is defined by assignement, can only be used inside of a function
	// UTF-8 uses one byte for ASCII characters and two and more bytes for unicode characters
	// "e" has one unicode point and "ě" has two unicode points
	interpretedString := "string is\n here"
	rawString := `string is\n here`
	fmt.Printf("\ninterpreted: %s, \nunicode byte: %v", interpretedString, []byte(interpretedString))
	fmt.Printf("\nraw: %s, \nunicode byte: %v\n", rawString, []byte(rawString))

	// Strings are immutable and can only be changed by assigning a new value
	// It doesn't allow you to create a pointer to any individual byte in a string
	// When you assign a new value both variable value and its pointer address will be updated to new one
	// Original string still exist in memory until it is no longer referenced by any variables,
	// at which point it will be eligible for garbage collection
	interpretedString = "hello👋"

	// len() returns number of bytes
	// UTF-8 is a variable encoding and some characters can take more than one byte
	fmt.Println(interpretedString, " ", len(interpretedString))

	// In Go string is just a slice of bytes
	// Slice is mutable
	bytesString := []byte(interpretedString)
	fmt.Println(len(bytesString), cap(bytesString))
	fmt.Println(string(bytesString))
	fmt.Println(bytesString)

	// When a string is converted to a run slice, the bytes stored in the string will be viewed as UTF-8 encoded byte sequence representations of many unicode code points
	// Bad UTF-8 encoding representations will be converted to a rune value for the unicode replacement character
	runesString := []rune(interpretedString)
	fmt.Println(len(runesString), cap(runesString))

	// You can convert byte slice to rune slice
	runesString = bytes.Runes(bytesString)
	fmt.Println(string(runesString))
	fmt.Println(runesString)

	// In these types of conversions value is not copied
}

var varInt int = 69

// When you pass a string variable to a function, Go creates a copy of the variable and passes it to the function.
// This is known as pass by value.
// Any changes made to the copy inside the function do not affect the original variable outside the function.
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
// Don't use pointers until we are sure we need them, it is only address to memory but it is stored in heap mememory so access to this is slower
// Param is type pointer type with int as base
func updateWithPointer(x *string) {
	*x = "James"
}

// This won't work becuase we are working with copy of value
func updateWithoutPointer(x string) {
	x = "John"
}

type person struct {
	name string
	age  int8
}

// In this case we are returning pointer address to value initialized in subsequent function, but it's useless becuase we cannot be sure what this points to in main function, so go creates copy of m var in heap, so the address stays the same but instead of stack the variable is saved in heap
// Avoid this at all costs because heap is a lot slower to access
// Running the program with -m gives you a way to analyze storage location of variables
func initPerson() *person {
	m := person{name: "jack", age: 12}
	fmt.Printf("%p", &m)
	return &m
}

// Only use pointers when:
// 1. We want to update value of variable passed in function
// 2. We're working with large data (File or big array) that are getting called a lot and we want optimize memory
func pointers() {
	i, j := 42, 2701

	// Pointer to address in memory of var i
	p := &i

	// Dereferncing value of var i
	fmt.Println(*p, i)
	fmt.Printf("%T\n", p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	name := "Jack"
	memoryName := &name

	// Every pointer is 8 bytes
	fmt.Println("\noriginal value:", name)
	fmt.Println("memory address:", memoryName)
	fmt.Println("value at memory address:", *memoryName)

	updateWithPointer(memoryName)

	fmt.Println("\nupdated value with pointer:", name)

	updateWithoutPointer(name)

	fmt.Println("\nupdated value without pointer:", name)

	newPerson := initPerson()

	fmt.Printf("\nInit person, %v, %p", newPerson, newPerson)

}

func runeType() {
	// Go alias for int32
	// It's intended to store unicode code point

	// Graphemes that fit in one unicode code can be specified within single quotes
	var data rune = '👍'
	fmt.Println(data)          //int32
	fmt.Printf("\n%c\n", data) // with format specifier

	// Graphemes that require more than one unicode code point cannot be stored in rune var
	var dataMultiPoint string = "👍🏿"
	fmt.Println(len([]rune(dataMultiPoint))) // 2 unicode code points
	fmt.Println(" ", len(dataMultiPoint))

	r := rune(0x1F60A) // Assigning a Unicode code point using type conversion
	fmt.Printf("\nrune2: %c\n", r)

	fmt.Printf("%c, %c, %c, %c, %c",
		97, // 97 is the code point for 'a'
		// Rune values
		'\141', // octal
		'\x61', // hex
		'\u0061',
		'\U00000061',
	)
}

// Prints
func DataTypes() {
	fmt.Println("\nData types: ")
	// %v is the default format, can be used with any type of variable
	// %d represents the value a decimal integer
	// %c is used for rune variables. It represents the value as the Unicode character corresponding to the given code point. The rune type holds a Unicode code point, and %c interprets that code point and displays the corresponding character.
	// %s is for strings
	// %f is for floating-point variables
	// %t is for boolean variables
	// %b is for binary representation of integers
	// %o is for octal representation of integers
	// %x is for hexadecimal representation of integers
	// %p is for pointer variables and represents value as a pointer address
	// %e is for scientific representation of floating-point variables
	// %T is for printing out name of variable type

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

	runeType()
}
