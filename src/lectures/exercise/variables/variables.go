//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var faveColor = "red"
	fmt.Println("favourit color", faveColor)
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	year, age := 1952, 71
	fmt.Println("born in", year, "age", age)
	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "D"
		lastInitial  = "H"
	)
	fmt.Println("initials", firstInitial, lastInitial)
	//* Declare (but don't assign!) a variable for your age (in days),
	var daysOld int
	//  then assign it on the next line by multiplying 365 with the age
	daysOld = age * 365
	// 	variable created earlier
	fmt.Println("i am", age, "days old", daysOld)

}
