package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in the GoLang")

	var fruitList [4] string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Mango"

	fmt.Println("The list of items in the array FruitList is: ", fruitList)
	fmt.Println("The number of items in the array FruitList is: ", len(fruitList))
	fmt.Println("The number of items in the array FruitList is: ", cap(fruitList))

	var newFruitList [10] string
	fmt.Println("The list of items in the array NewFruitList is: ", newFruitList)
	fmt.Println("The number of items in the array NewFruitList is: ",len(newFruitList))
	fmt.Println("The number of items in the array NewFruitList is: ",cap(newFruitList))

	// If the size of the array is not known at compile time, then it grow automatically.
	var it=[]string{"A"}
	fmt.Println("The list of items in the array It is: ", it)
	fmt.Println("The number of items in the array It is: ",len(it))

	// If we declare the size of the array as 10 , the length of the array will be 2 only for extra we have error
	var iti=[2]string{"A","B","C"}
	fmt.Println("The list of items in the array It is: ", iti)
	fmt.Println("The number of items in the array It is: ",len(iti))
}
