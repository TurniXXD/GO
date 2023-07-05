package functions

// First-class function => function that is treated like any other variable
// Higher order function => function that takes another function as an argument, may be used in http handlers
// Function currying => function takes a new function as an argument, may be used in middleware

import (
	"fmt"
	"math"
	"math/big"
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

// If we want to update receiver value declare it as pointer receiver, otherwise value wouldn't be updated
func (c *circle) setDefaultRadius() {
	c.radius = 4.0
}

func printShapeInfo(s shape) {
	fmt.Printf("area of shape %T is: %0.2f \n", s, s.area())

	fmt.Println()
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

type user struct {
	name    string
	isAdmin bool
}

func deferKeyword(users map[string]user, name string) (log string) {
	// Defer creates operation that will happen right before the return of a function
	src, err := os.Open("someFile")
	if err != nil {
		return
	}

	// It is mostly used for cleanup and closing instances so we don't waste resources
	defer src.Close()

	defer delete(users, name)

	user, ok := users[name]
	if !ok {
		// instead of specifying this before every return
		// delete(users, name)
		return "not found"
	}
	if user.isAdmin {
		// delete(users, name)
		return "admin deleted"
	}
	// delete(users, name)
	return "user deleted"
}

// Closure function references variables from outside its own function body
// The function may access and assign to the referenced variables
func concatter() func(string) string {
	// When we call concatter we're saving reference to this doc variable
	doc := ""
	// Every subsequent call to this function will keep adding on words into doc
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func emailSender() func(string, string) int {
	userCount := make(map[string]int)

	return func(user, content string) int {
		fmt.Printf("send email '%s' to user %s, ", content, user)
		userCount[user]++
		return userCount[user]
	}

}

// Closure function can mutate a variable outside its own function body
// Function then has acces to a mutable reference to the original value
func closureFunc() {
	wordAggregator := concatter()

	wordAggregator("Welcome")
	wordAggregator("Jacob")
	fmt.Println(wordAggregator("here!"))

	emailCounter := emailSender()
	emailCounter("Denis", "Hello worker")
	emailCounter("Bob", "Hello man")
	emailCounter("James", "Hello firefighter")
	emailCounter("James", "How are you firefighter")
	fmt.Println(emailCounter("James", "Hello doctor"))
}

func FibInt(n int) int {
	if n < 2 {
		return n
	}

	a, b := 0, 1
	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}

func FibBig(n uint64) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	a, b := big.NewInt(0), big.NewInt(1)
	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}

	return b
}

// These functions are referenced in native_functions_test.go file
func NativeFunctions() {
	fmt.Println(FibInt(30)) // 832040
	fmt.Println(FibBig(30)) // 832040
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
	circle1.setDefaultRadius()
	printShapeInfo(circle1)

	curriedFunc()

	users := make(map[string]user)
	users["jerry"] = user{
		name:    "jerry",
		isAdmin: false,
	}

	users["daniel"] = user{
		name:    "daniel",
		isAdmin: true,
	}

	deferKeyword(users, "jerry")
	deferKeyword(users, "daniel")

	closureFunc()

	NativeFunctions()
}
