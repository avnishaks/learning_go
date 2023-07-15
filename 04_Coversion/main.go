package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Our Pizza App")
	fmt.Println("Please rate the pizza between 1 to 10")
	reader := bufio.NewReader(os.Stdin)
	output, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating the pizza", output)
	fmt.Printf("Type of the output value is %T\n", output)

	numRatings, err := strconv.ParseInt(strings.TrimSpace(output),10,64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRatings+1)
	}
}
