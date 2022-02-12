package math

import (
	"fmt"
	"math"
)

func Math() {
	fmt.Println("\nMath: ")
	//use float in order to get result in float
	fmt.Println(5 / 2)
	fmt.Println(5.0 / 2.0)
	fmt.Println(5 / float64(2))

	fmt.Println(math.Log(10))
	fmt.Println(math.Log2(16))

	fmt.Println(math.Pow(2, 4))
	fmt.Println(math.Pow10(4))
}
