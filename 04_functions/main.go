package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int { // <- shorthand getSum(num1 int, num2 int)
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Zay"))
	fmt.Println(getSum(3, 4))
}
