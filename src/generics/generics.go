package generics

import "fmt"

func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitSlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

type stringer interface {
	String() string
}

type stringerType struct {
	Data string
}

func (st stringerType) String() string {
	return st.Data
}

func concat[T stringer](values []T) string {
	result := ""
	for _, v := range values {
		result += v.String()
	}
	return result
}

// Parametric constraint
type equalizer[T any] interface {
	Equal(other T) bool
}

func HasWithInterface[T equalizer[T]](list []T, value T) bool {
	for _, v := range list {
		if value.Equal(v) {
			return true
		}
	}

	return false
}

type ID int

func (id ID) Equal(other ID) bool {
	return id == other
}

func Generics() {
	fmt.Println("\nGenerics:\n ")

	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"a", "b", "c", "d", "e", "f"}
	stringerSlice := []stringer{
		stringerType{Data: "Hello "},
		stringerType{Data: "World!"},
	}

	// Instead of using separate functions for different types
	fmt.Println(splitIntSlice(intSlice))
	fmt.Println(splitStringSlice(stringSlice))

	// We can just use generic
	fmt.Println(splitSlice[int](intSlice))
	fmt.Println(splitSlice[string](stringSlice))

	fmt.Println(concat(stringerSlice))

	fmt.Println(min[int16](1, 4))
	fmt.Println(min[int8](2, 1))

	fmt.Printf("Has 1 %v\n", HasWithInterface([]ID{1, 2}, 1))
}
