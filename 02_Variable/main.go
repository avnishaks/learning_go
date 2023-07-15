package main

import "fmt"

func main() {
	// type is of string type
	var username string = "Avnish Kumar Singh"
	fmt.Println(username)
	fmt.Printf("Variable type: %T\n", username)

	// type is of bool type
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type: %T\n", isLoggedIn)

	// type is of int type
	var age uint8 = 255
    fmt.Println(age)
    fmt.Printf("Variable type: %T\n", age)

	// default value for the string , bool and int
	var life int 
	var name string
	var isDone bool
	// default value for the int is 0
	fmt.Println(life)
	// default value for the string is ""
	fmt.Println(name)
	// default value for the bool is false
	fmt.Println(isDone)


	// Implicit type 
	var Love="Love in the Air"
	fmt.Println(Love)

	// no var style -> [ := -> this type of the operator is allowed inside the function block only. ] 
	numbercount := 10
	fmt.Println(numbercount)

}
