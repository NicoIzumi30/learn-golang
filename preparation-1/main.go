package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	fmt.Println("Hello World")
// }

func main() {
	var namaVariable1 string = "Hello World " //Long declaration
	namaVariable2 := "Hello World 2"          //Short declaration

	fmt.Println(namaVariable1)
	fmt.Println(namaVariable2)

	// Primitive : boolean, int, float, string

	// Boolean
	boolVar := true
	fmt.Printf("Type : %T Value : %v\n", boolVar, boolVar)

	// Integer
	intVar := int(10)
	intVarVar1 := int32(12)
	intVarVar2 := int64(13)
	fmt.Printf("Type : %T Value : %v\n", intVar, intVar)
	fmt.Printf("Type : %T Value : %v\n", intVarVar1, intVarVar1)
	fmt.Printf("Type : %T Value : %v\n", intVarVar2, intVarVar2)

	// Float
	floatVar := float32(10.5)
	floatVarVar1 := float64(12.5)
	fmt.Printf("Type : %T Value : %v\n", floatVar, floatVar)
	fmt.Printf("Type : %T Value : %v\n", floatVarVar1, floatVarVar1)

	// Byte
	byteVar := byte(255)
	fmt.Printf("Type : %T Value : %v\n", byteVar, byteVar)
	byteVar2 := []byte("Hello World")
	fmt.Printf("Type : %T Value : %v\n", byteVar2, byteVar2)

	// Rune
	runeVar := rune('O')
	fmt.Printf("Type : %T Value : %v\n", runeVar, runeVar)

	// Complex
	complexVar := -7 + 3i
	fmt.Printf("Type : %T Value : %v\n", complexVar, complexVar)

	// Error
	errorVar := errors.New("Error detected")
	fmt.Printf("Type : %T Value : %v\n", errorVar, errorVar)

	// Interface
	var myInterfaceVar interface{}
	myInterfaceVar = 5
	fmt.Printf("Type : %T Value : %v\n", myInterfaceVar, myInterfaceVar)
	myInterfaceVar = "Hello World"
	fmt.Printf("Type : %T Value : %v\n", myInterfaceVar, myInterfaceVar)
}

// Interface
type MethodList interface {
	myFunction()
	myFunction2(int) int
}
