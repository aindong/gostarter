package main

import "fmt"

func main() {
	// Declaring Variables
	var name = "Alleo"
	const age = 21
	const isCool = true

	// Different ways of printing to console
	fmt.Println("Hello", name, ", good luck on learning GO!")
	fmt.Println("Your age is", age)
	fmt.Printf("You are cool: %v \n", isCool)

	var userName string = "Leo"
	fmt.Println("Hello", userName, ", good luck on learning GO!")

	// Operations
	var x int = 1
	var y int = 2
	fmt.Println("Sum of x and y is: ", x+y)
	fmt.Println("Product of x and y is: ", x*y)
	fmt.Println("Difference of x and y is: ", x-y)
	fmt.Println("Division of x and y is: ", x/y)
	fmt.Println("Remainder of x and y is: ", x%y)

	// Float data type
	var a float32 = 1.3
	var b float32 = 2.2
	fmt.Println("Sum of a and b is: ", a+b)
	fmt.Println("Product of a and b is: ", a*b)

	// Getting user input from cli
	var input string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&input)

	var inputAge int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&inputAge)

	// Conditions
	if inputAge < 18 {
		fmt.Println("You are not old enough to enter the club")
		return
	}

	switch inputAge {
	case 18:
		fmt.Println("You are 18 years old")
	case 19:
		fmt.Println("You are 19 years old")
	case 20:
		fmt.Println("You are 20 years old")
	default:
		fmt.Println("You are not old enough to enter the club")
		return
	}

	fmt.Println("Hello", input, ", good luck on learning GO!")
}
