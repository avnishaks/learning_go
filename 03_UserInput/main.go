package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "Welcome to User Input Manager"
	fmt.Println(Welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	// comma ok || err 
	input, _ := reader.ReadString('\n')
	fmt.Println("Hey , their your entered name is: ",input)
	fmt.Printf("Type of entered name is: %T",input)
}
