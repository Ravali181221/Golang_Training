package main

import "fmt"

// Declare the variable x
// ASSIGN the value 21
// declare and assign = initialization
var x = 21

// DECLARE there is a varaible with the INDENTIFIER z
// and that the variable with the INDENTIFIER z is of TYPE int
// ASSIGN the ZERO VALUE of TYPE int to z
// false for boolean, 0 for integers, 0.0 for floats, "" for strings.
// and nill for pointers, functions, interfaces, slices, channels and maps
var z int

func main1() {
	//declare a variable and assign the value
	foovar()
	x := 42 //Initialization and assignment operation

	fmt.Println(z)

	fmt.Println(x)
	x = 99 //It's only Assignment operation..it's like reassignment
	fmt.Println(x)
	y := 100 + 24 //statement made up of expression(operand(numbers,strings,etc) operator(+,-,etc))
	//expression is made up of explicit values
	fmt.Println(y)
	//gophers are used for logo's for websites

	z := "Ravali"
	fmt.Println(z)

}

func foovar() {
	fmt.Println(x)
}
