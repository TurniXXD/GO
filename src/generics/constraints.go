package generics

import "fmt"

type customConstraint interface {
	// Union of types that can be in constraint
	// ~ operator means that all types the have underlying string type can also be in this contstraint
	int | float64 | ~string
}

// This can also be in the interface
type customString string

// We can specify allowed type of generic
func min[T int8 | int16 | int32](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Comparable constraint accepts any type that implements the equality operator "==" or "!="
type Set[T comparable] map[T]struct{}

type SetCustom[T customConstraint] map[T]struct{}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}

func (s SetCustom[T]) HasCustom(value T) bool {
	_, ok := s[value]
	return ok
}

func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

func NewCustomSet[T customConstraint](values ...T) SetCustom[T] {
	set := make(SetCustom[T], len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

func HasAny[T any](list []T, value T, equal func(a, b T) bool) bool {
	for _, v := range list {
		// any type has uncomparable types in type set
		// if v == value {
		// 	return true
		// }

		if equal(v, value) {
			return true
		}
	}
	return false
}

func Constraints() {
	intSet := NewSet(1, 2, 3)
	fmt.Printf("Has 2 %v\n", intSet.Has(2))
	fmt.Printf("Has 5 %v\n", intSet.Has(5))

	stringSet := NewSet("a", "b", "c", "d", "e", "f")
	fmt.Printf("Has a %v\n", stringSet.Has("a"))
	fmt.Printf("Has i %v\n", stringSet.Has("i"))

	// Custom constraint
	customStringSet := NewCustomSet[customString]("a", "b", "c", "d", "e", "f")
	fmt.Printf("Has f %v\n", customStringSet.HasCustom("f"))

	equalInt := func(a, b int) bool {
		return a == b
	}

	fmt.Printf("Has 5 %v\n", HasAny([]int{1, 2}, 5, equalInt))
	fmt.Printf("Has 2 %v\n", HasAny([]int{1, 2}, 2, equalInt))
}
