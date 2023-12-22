package main

import "fmt"

// main function
func main() {
	// // B1---------------------------------------------
	// fmt.Printf("Hello World!")

	// // B2---------------------------------------------
	// fmt.Println("Enter Your First Name: ")
	// // var then variable name then variable type
	// var first string
	// // Taking input from user
	// fmt.Scanln(&first)
	// fmt.Println("Enter Second Last Name: ")
	// var second string
	// fmt.Scanln(&second)

	// // Print function is used to
	// // display output in the same line
	// fmt.Print("Your Full Name is: ")

	// // Addition of two string5
	// fmt.Println(first + " " + second)

	// // B3---------------------------------------------
	// fmt.Println("Enter a number you want: ")
	// var number int
	// fmt.Scanln(&number)
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("%d x %d = %d\n", number, i, number*i)
	// }
	
	// B4---------------------------------------------
	fmt.Println("Enter your array: ")
	var array [10]int
	for i := 0; i <= range(array); i++ {
		var temp int
		fmt.Scanln(&temp)
		array = append(array, temp)
	}
}
