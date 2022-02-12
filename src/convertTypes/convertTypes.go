package convertTypes

import (
	"fmt"
	"reflect"
	"strconv"
)

func strToInt() {
	// Atoi => ASCII to int
	intVar, _ := strconv.Atoi("666")
	fmt.Println(intVar, reflect.TypeOf(intVar))
}

func strToIntBase() {
	intVarInRange, errIn := strconv.ParseInt("32767", 0, 16)
	intVarOutOfRange, errOut := strconv.ParseInt("32768", 0, 16)
	fmt.Println(intVarInRange, errIn, reflect.TypeOf(intVarInRange))
	fmt.Println(intVarOutOfRange, errOut, reflect.TypeOf(intVarOutOfRange))
}

func scannedStrToInt() {
	intValue := 0
	// Scans string arg and stores it into var
	_, err := fmt.Sscan("666", &intValue)
	fmt.Println(intValue, err, reflect.TypeOf(intValue))
}

func strToFloat() {
	f, err := strconv.ParseFloat("3.1415926535", 32)
	f32 := float32(f)
	fmt.Println(f, err, reflect.TypeOf(f))
	fmt.Println(f32, err, reflect.TypeOf(f32))
}

func ConvertTypes() {
	fmt.Println("\nConvert types: ")
	strToInt()
	strToIntBase()
	scannedStrToInt()
	strToFloat()
}
