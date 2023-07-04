package functions

// First-class function => function that is treated like any other variable
// Higher order function => function that takes another function as an argument, may be used in http handlers
// Function currying => function takes a new function as an argument, may be used in middleware

import (
	"fmt"
	"math"
	"os"
	"reflect"
)

func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func addReturn(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func addReturnAlias(x int, y int) (summary int) {
	summary = x + y
	return
}

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return
}

func passingAddress(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
}

var (
	anonymousFunc = func(l int, b int) int {
		return l * b
	}
)

func sum(x, y int) int {
	return x + y
}

func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

// pass variable num of args
func variadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

func variadicExampleWithNormArg(a int, s ...int) {
	fmt.Println(a * s[0])
	fmt.Println(a * s[3])
}

func variadicExampleWithInterface(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

// Schedule a function call with defer
func deferFunc() bool {
	fmt.Println("Loading file...")
	defer fmt.Println("File loading finished")
	csvFile, err := os.Open("./test/test.csv")

	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("File loaded successfully")
	defer csvFile.Close()
	return true
}

func deferredListFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// Interface groups types together based on their methods
// For type to be classified as interface it needs to have all methods of an interface
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

// Now every circle var has this function
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("area of shape %T is: %0.2f \n", s, s.area())
}

// Curried function
// someName takes a function someFunc that takes two int arguments and returns an int.
func someName(someFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return someFunc(x, 2)
	}
}

func addFunc(a, b int) int {
	return a + b
}

func curriedFunc() {
	// Creating a curried function using someName and the add function as an argument.

	// Calling the curried function addTwo with an argument.
	newFunc := someName(addFunc)
	result := newFunc(3)

	fmt.Println("Curried function result:", result) // Output: 5
}

// To export a func first letter of a func needs to be a capital letter
func Functions() {
	fmt.Println("\nFunctions: ")
	add(5, 3)

	fmt.Println(addReturn(20, 30))
	fmt.Println(addReturnAlias(1, 3))

	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)

	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)
	// get variable's address in memory
	ageAddr := &age
	passingAddress(ageAddr, &text)
	fmt.Println("After :", text, age)

	fmt.Println(anonymousFunc(20, 30))
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
	l := 20
	b := 30
	// Anonymous function with closures can access variables defined outside body.
	func() {
		area := l * b
		fmt.Println(area)
	}()

	partial := partialSum(3)
	fmt.Println(partial(7))

	fmt.Println(squareSum(5)(6)(7))

	variadicExample("red", "blue", "green", "yellow")
	variadicExampleWithNormArg(1, 5, 7, 8, 10)
	variadicExampleWithInterface(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})

	deferFunc()
	deferredListFunc()

	circle1 := circle{radius: 5.0}
	fmt.Printf("area of circle with radius: %0.2f is %0.2f \n", circle1.radius, circle1.area())
	printShapeInfo(circle1)

	curriedFunc()
}
