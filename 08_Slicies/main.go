package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hey ! Geek Welcome to Golang")

	var fruitList = []string{"Apple", "Banana", "Orange"}

	fmt.Printf("Type of fruitList: %T\n", fruitList)

	fmt.Println("List of Fruit: ", fruitList)

	fruitList = append(fruitList, "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M")

	fmt.Println("List of Fruit AFTER APPENDING THE VALUE : ", fruitList)

	// Slicing Concept in Golang
	fruitList = append(fruitList[1:2])
	fmt.Println(fruitList)

	// appned alloate the new memeory to append the item in the list

	highscores := make([]int, 2)
	highscores[0] = 100
	highscores[1] = 200
	//highscores[2]=300; // Gives error as static memory allocation happen for highscores

	fmt.Println("The Highscores are: ", highscores)

	// Append create memory dynamically for highscores
	highscores = append(highscores, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 111)
	fmt.Println("Highscores value after updation is : ", highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	// How to remove an element from the slices

	var courses = []string{"Java", "C++", "Python", "Go", "JavaScript", "react", "native"}
	fmt.Println(courses)
	var index int= 1 // 1 index element C++ will be removed
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	

}
