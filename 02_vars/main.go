package main

import "fmt"

func main() {
	// Main TYPE
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Zaycho"
	var age int32 = 29
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	// name := "Zaycho"
	// email := "zaycho_89@yahoo.com"

	name, email := "Zaycho", "zaycho_89@yahoo.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
