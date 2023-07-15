package main

import "fmt"

func main() {
	fmt.Println("Hey ! Geek Welcome in the world of the Golang Pointers")
	myNumber := 10
	var ptr=&myNumber
	fmt.Println("The value of the declared variable is :", *ptr);
	fmt.Println("The address of the declared variable is :", ptr);
	// pointer gurantees that what the operation is perfomed is perfomed on the actual memory location not the copy of the variable

	*ptr=*ptr*2+10;
	fmt.Println("The updated value of the declared variable is :", *ptr);
}
