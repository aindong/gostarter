package main

import (
	"fmt"
	"sync"
	"time"
)

// var name string = "Leo"

type User struct {
	firstName string
	lastName  string
	age       uint
}

var wg = sync.WaitGroup{}

func main() {
	// Declaring Variables
	// var name = "Alleo"
	// const age = 21
	// const isCool = true

	// // Different ways of printing to console
	// fmt.Println("Hello", name, ", good luck on learning GO!")
	// fmt.Println("Your age is", age)
	// fmt.Printf("You are cool: %v \n", isCool)

	// var userName string = "Leo"
	// fmt.Println("Hello", userName, ", good luck on learning GO!")

	// // Operations
	// var x int = 1
	// var y int = 2
	// fmt.Println("Sum of x and y is: ", x+y)
	// fmt.Println("Product of x and y is: ", x*y)
	// fmt.Println("Difference of x and y is: ", x-y)
	// fmt.Println("Division of x and y is: ", x/y)
	// fmt.Println("Remainder of x and y is: ", x%y)

	// // Float data type
	// var a float32 = 1.3
	// var b float32 = 2.2
	// fmt.Println("Sum of a and b is: ", a+b)
	// fmt.Println("Product of a and b is: ", a*b)

	// // Getting user input from cli
	// var input string
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&input)

	// var inputAge int
	// fmt.Print("Enter your age: ")
	// fmt.Scanln(&inputAge)

	// // Conditions
	// if inputAge < 18 {
	// 	fmt.Println("You are not old enough to enter the club")
	// 	return
	// }

	// switch inputAge {
	// case 18:
	// 	fmt.Println("You are 18 years old")
	// case 19:
	// 	fmt.Println("You are 19 years old")
	// case 20:
	// 	fmt.Println("You are 20 years old")
	// default:
	// 	fmt.Println("You are not old enough to enter the club")
	// 	return
	// }

	// fmt.Println("Hello", input, ", good luck on learning GO!")

	// Arrays
	// var friends [3]string
	// friends[0] = "foo"
	// friends[1] = "bar"
	// friends[2] = "baz"
	// fmt.Println(friends) // Priting array

	// // another way to declare an array
	// var friends2 = [3]string{"foo", "bar", "baz"}
	// fmt.Println(friends2) // Priting array

	// // print size of array
	// const totalFriends = len(friends) + len(friends2)
	// fmt.Println("You have", totalFriends, "friends")

	// Slices
	// var futureFriends []string
	// futureFriends = append(futureFriends, "foo1")
	// futureFriends = append(futureFriends, "foo2")

	// fmt.Println(futureFriends)

	// // Loops
	// for i := 0; i < len(futureFriends); i++ {
	// 	fmt.Println(futureFriends[i])
	// }

	// // iterate over a slice
	// // blank identifier "_" is used to ignore the value
	// for i, friend := range futureFriends {
	// 	fmt.Println(i, friend)
	// }

	// // Count to 10
	// for i := 0; i <= 10; i++ {
	// 	fmt.Print(i)

	// 	if i >= 0 && i < 10 {
	// 		fmt.Print(" ")
	// 	}
	// }

	// multiple := getTheMultipleOf(2, 3)
	// fmt.Println(multiple)

	// printName()
	// sayQuote()

	// fmt.Println("Using the utils package:", utils.Add(2, 3))

	// // Maps
	// var user = make(map[string]string)
	// user["firstName"] = "Leo"
	// user["lastName"] = "Tech"

	// fmt.Println(user)

	// // list of users
	// var users = make([]map[string]string, 0)
	// users = append(users, user)

	// var user2 = make(map[string]string)
	// user2["firstName"] = "Leo 2"
	// user2["lastName"] = "Tech 2"
	// users = append(users, user2)

	// for _, user := range users {
	// 	fmt.Println(user["firstName"])
	// }

	var users = make([]User, 0)
	users = append(users, User{"Leo", "Tech", 21})
	users = append(users, User{"Leo 2", "Tech 2", 22})

	var user3 = User{
		firstName: "Leo 3",
		lastName:  "Tech 3",
		age:       23,
	}
	users = append(users, user3)

	for _, user := range users {
		fmt.Println(user.firstName)
	}

	wg.Add(1)
	go sendEmail()

	wg.Wait()
}

func sendEmail() {
	fmt.Println("Sending email")
	// Delay the function
	time.Sleep(time.Second * 10)

	fmt.Println("Email sent")
	wg.Done()
}

// func getTheMultipleOf(x int, y int) int {
// 	return x * y
// }

// func printName() {
// 	fmt.Println(name)
// }
