package operations

import (
	"fmt"
	"math/rand"
	"strings"
)

func Array() {
	fmt.Println("Array operation")

	/*

		[] Arrays
		- Fixed length
		- Same type
		- Indexable
		- Contiguous in Memory

	*/

	var intArr []int32 = []int32{1, 2, 3, 4, 6, 6, 6, 77, 775}
	fmt.Println(intArr[0])
	fmt.Println(intArr[7:9])

	var sliceintArr []int32 = []int32{2, 3, 4, 4, 5, 664, 4223, 443}
	intArr = append(intArr, sliceintArr...)
	fmt.Println(intArr)

}

func Map() {

	fmt.Println("Map operation in Go")

	// creating a map in go
	var myMap map[int32]int32 = make(map[int32]int32)

	fmt.Println(myMap)

	// assigning the value to the map in go
	var myMap2 = map[int32]int32{11: 121, 12: 144}
	fmt.Println(myMap2[12])

	// creating a map for string -> string
	var myMap3 = map[string]string{"deep": "Shikha", "avnish": "kumar"}
	fmt.Println(myMap3["deep"])
	fmt.Println(myMap3["avnish"])
	delete(myMap3, "deep")
	var mapValue, isPresent = myMap3["deep"]
	fmt.Println(isPresent)
	fmt.Println(mapValue)

	// That's all about map in go all point is covered
}

func Loop() {
	var myMap3 = map[string]string{"deep": "Shikha", "avnish": "kumar"}
	for name, age := range myMap3 {
		fmt.Printf("Name: %v and SurName : %v\n", name, age)
	}

	// printing the number form 0 to 10

	for i := 0; i < 10; i++ {
		fmt.Print(i, " | ")
	}

	fmt.Println("\n")

	fmt.Println(rand.Intn(200))
}

func StringRunesByte() {
	var myString = []rune("resume")
	var indexed = myString[1]
	fmt.Printf("%v , %T \n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of the string is %v\n", len(myString))

	var strSlice = []string{"a", "b", "c", "d", "e", "f", "g"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var dummyString = strBuilder.String()
	fmt.Printf("%v\n", dummyString)
}
